package main

import (
	"testing"
)

type LevenshteinCase struct {
	Source      string
	Destination string
	Distance    int
}

func Test_LevensteinDistance(test *testing.T) {
	var cases []LevenshteinCase = []LevenshteinCase{
		LevenshteinCase{"sitting", "kitten", 3},
		LevenshteinCase{"sunday", "saturday", 3},
		LevenshteinCase{"", "foobar", len("foobar")},
		LevenshteinCase{"foobar", "", len("foobar")},
		LevenshteinCase{"bingis", "bingis", 0},
	}

	var distance int
	var testcase LevenshteinCase
	for _, testcase = range cases {
		if distance = LevenshteinDistance(testcase.Source, testcase.Destination); distance != testcase.Distance {
			test.Errorf(
				"%s -> %s distance is %d, not %d",
				testcase.Source,
				testcase.Destination,
				distance,
				testcase.Distance,
			)
		}
	}
}
