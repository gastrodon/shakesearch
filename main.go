package main

import (
	"github.com/gastrodon/groudon/v2"

	"net/http"
)

const (
	ROUTE_SEARCH = "^/search/?$"
)

func main() {
	var err error
	var searcher Searcher = Searcher{}
	if err = searcher.Load("completeworks.txt"); err != nil {
		panic(err)
	}

	groudon.AddHandler("GET", ROUTE_SEARCH, handleSearch(searcher))
	http.HandleFunc("/", groudon.Route)
	http.ListenAndServe(":3001", nil)
}

func handleSearch(searcher Searcher) func(*http.Request) (int, map[string]interface{}, error) {
	return func(request *http.Request) (code int, RMap map[string]interface{}, err error) {
		var query []string
		var exists bool
		if query, exists = request.URL.Query()["q"]; !exists || len(query) != 1 {
			code = 400
			return
		}

		code = 200
		RMap = map[string]interface{}{
			"results": searcher.Search(query[0]),
		}
		return
	}
}
