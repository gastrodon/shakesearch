package main

import (
	"index/suffixarray"
	"io/ioutil"
	"regexp"
	"strings"
)

const (
	RESULT_LIMIT = 128
)

var (
	nonLetter *regexp.Regexp = regexp.MustCompile(`[^a-z0-9]+`)
)

func min(candidates ...int) (it int) {
	it = candidates[0]

	var candidate int
	for _, candidate = range candidates[1:] {
		if candidate < it {
			it = candidate
		}
	}

	return
}

func max(candidates ...int) (it int) {
	it = candidates[0]

	var candidate int
	for _, candidate = range candidates[1:] {
		if candidate > it {
			it = candidate
		}
	}

	return
}

type Wordset struct {
	Word      string
	Frequency int
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
		field = nonLetter.ReplaceAllString(field, "")
		if _, exists = search.words[field]; !exists {
			search.words[field] = 1
		} else {
			search.words[field]++
		}
	}

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

func (search *Searcher) WordsOrdered() (set []Wordset) {
	set = make([]Wordset, len(search.words))

	var index int = 0
	var word string
	var frequency int
	for word, frequency = range search.words {
		set[index] = Wordset{word, frequency}
		index++
	}

	for index = range set[:len(set)-1] {
		for set[index].Frequency > set[index+1].Frequency {
			if index == 0 {
				break
			}

			set[index+1], set[index] = set[index], set[index+1]
			index--
		}
	}

	return
}
