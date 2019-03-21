package main

import (
	"crawler/fronted/controller"
	"net/http"
)

func main() {
	http.Handle("/search", controller.CreateSearchResultHandler("fronted/view/template.html"))

	http.Handle("/", http.FileServer(http.Dir("fronted/view")))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
