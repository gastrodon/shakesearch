package main

import (
	"fmt"
	"index/suffixarray"
	"io/ioutil"
	"strings"
)

const (
	RESULT_LIMIT = 128
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

type Searcher struct {
	source      string
	suffixArray *suffixarray.Index
	size        int
	words       map[string]int
}

func (search *Searcher) Load(file string) (err error) {
	var data []byte
	if data, err = ioutil.ReadFile(file); err != nil {
		return
	}

	search.source = string(data)
	search.suffixArray = suffixarray.New(data)
	search.size = len(data)

	var fields []string = strings.Fields(strings.ToLower(search.source))
	search.words = make(map[string]int, len(fields))

	var exists bool
	var field string
	for _, field = range fields {
		if _, exists = search.words[field]; !exists {
			search.words[field] = 1
		} else {
			search.words[field]++
		}
	}

	fmt.Println(len(fields))
	fmt.Println(len(search.words))

	return
}

func (search *Searcher) Search(query string) (results []string) {
	var indexes []int = search.suffixArray.Lookup([]byte(query), RESULT_LIMIT)
	results = make([]string, len(indexes))

	var index, head, tail int
	for index = range indexes {
		head = max(0, indexes[index]-250)
		tail = min(search.size, indexes[index]+250)
		results[index] = search.source[head:tail]
	}

	return
}
