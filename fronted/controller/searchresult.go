package controller

import (
	"context"
	"crawler/engine"
	"crawler/fronted/model"
	"crawler/fronted/view"
	"gopkg.in/olivere/elastic.v6"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) (resultHandler SearchResultHandler) {
	newClient, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: newClient,
	}
}

func (h SearchResultHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	var page model.SearchResult
	page, err = h.getSearchResult(q, from)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(resp, page)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (searchResult model.SearchResult, err error) {
	var result model.SearchResult
	result.Query = q
	resp, err := h.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())
	if err != nil {
		return result, err
	}
	result.Hints = resp.TotalHits()
	result.Start = from
	for _, value := range resp.Each(reflect.TypeOf(engine.Item{})) {
		result.Items = append(result.Items, value.(engine.Item))
	}
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}
