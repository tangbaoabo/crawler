package parser

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestParseCityList(t *testing.T) {

	contents, _ := ioutil.ReadFile("citylist.html")
	parseResult := ParseCityList(contents)
	for _, value := range parseResult.Items {
		log.Printf("the result is: %s\n", value)
	}
}
