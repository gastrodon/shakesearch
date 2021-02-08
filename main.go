package main

import (
	"github.com/gastrodon/groudon/v2"

	"fmt"
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

func (s *Searcher) Load(filename string) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Load: %w", err)
	}
	s.CompleteWorks = string(dat)
	s.SuffixArray = suffixarray.New(dat)
	s.size = len(dat)
	return nil
}

func (s *Searcher) Search(query string) []string {
	idxs := s.SuffixArray.Lookup([]byte(query), -1)
	results := []string{}
	for _, idx := range idxs {
		results = append(results, s.CompleteWorks[max(0, idx-250):min(s.size, idx+250)])
	}
	return results
}
