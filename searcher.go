package main

import (
	"index/suffixarray"
	"io/ioutil"
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
	CompleteWorks string
	SuffixArray   *suffixarray.Index
	size          int
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
	var indexes []int = search.SuffixArray.Lookup([]byte(query), RESULT_LIMIT)
	results = make([]string, len(indexes))

	var index, head, tail int
	for index = range indexes {
		head = max(0, indexes[index]-250)
		tail = min(search.size, indexes[index]+250)
		results[index] = search.CompleteWorks[head:tail]
	}

	return
}
