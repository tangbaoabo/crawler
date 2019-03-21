package view

import (
	"crawler/fronted/model"
	"html/template"
	"io"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(fileName string) SearchResultView {
	return SearchResultView{
		template: template.Must(template.ParseFiles(fileName)),
	}
}
func (s SearchResultView) Render(writer io.Writer, data model.SearchResult) (err error) {
	return s.template.Execute(writer, data)
}
