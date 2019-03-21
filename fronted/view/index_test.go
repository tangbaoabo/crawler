package view

import (
	"crawler/domain"
	"crawler/engine"
	"crawler/fronted/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {

	view := CreateSearchResultView("template.html")
	file, _ := os.Create("index.test.html")

	page := model.SearchResult{}
	page.Hints = 123
	profile := domain.Profile{
		Name:      "老猫",
		Gender:    "男士",
		Age:       56,
		Height:    174,
		Weight:    90,
		WorkPlace: "阿里日土",
		InCome:    "2-5万",
		Marriage:  "离异",
		HoKou:     "北京",
		XinZuo:    "天蝎座(10.23-11.21)",
		House:     "已购房",
		Car:       "已买车",
	}
	item := engine.Item{
		Url:     "https://localhost:9090/tangbaobao",
		Type:    "zhenai",
		Id:      "tangbaobao",
		PayLoad: profile,
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err := view.Render(file, page)

	if err != nil {
		panic(err)
	}
}
