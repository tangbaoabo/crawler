package persist

import (
	"context"
	"crawler/engine"
	"crawler/domain"
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v6"
	"testing"
)

func TestSaver(t *testing.T) {
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

	expected := engine.Item{
		Url:     "https://localhost:9090/tangbaobao",
		Type:    "zhenai",
		Id:      "tangbaobao",
		PayLoad: profile,
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	var index = "dating_test"

	err = save(expected, client, index)

	if err != nil {
		panic(err)
	}

	result, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	var actual engine.Item
	err = json.Unmarshal(*result.Source, &actual)
	if err != nil {
		panic(err)
	}
	actualProfile, err := domain.FromJsonObj(actual.PayLoad)
	actual.PayLoad = actualProfile

	if expected != actual {
		t.Errorf("got a error %v\n not equals %v", expected, actual)
	}

	fmt.Println(expected)

	fmt.Println()

	fmt.Println(actual)

}
