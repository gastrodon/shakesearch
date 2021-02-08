package main

import (
	"github.com/gastrodon/groudon/v2"

	"index/suffixarray"
	"io/ioutil"
	"net/http"
)

const (
	ROUTE_SEARCH = "^/search/?$"
)

func max(this, that int) (it int) {
	if this < that {
		it = that
		return
	}

	it = this
	return
}

func min(this, that int) (it int) {
	if this > that {
		it = that
		return
	}

	it = this
	return
}

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

type Searcher struct {
	CompleteWorks string
	SuffixArray   *suffixarray.Index
	size          int
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

func (search *Searcher) Load(file string) (err error) {
	var data []byte
	if data, err = ioutil.ReadFile(file); err != nil {
		return
	}

	search.CompleteWorks = string(data)
	search.SuffixArray = suffixarray.New(data)
	search.size = len(data)
	return
}

func (search *Searcher) Search(query string) (results []string) {
	var indexes []int = search.SuffixArray.Lookup([]byte(query), -1)
	results = make([]string, len(indexes))

	var index, head, tail int
	for index = range indexes {
		head = max(0, indexes[index]-250)
		tail = min(search.size, indexes[index]+250)
		results[index] = search.CompleteWorks[head:tail]
	}

	return
}
