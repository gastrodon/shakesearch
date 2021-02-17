package main

import (
	"testing"
)

type CompareCase struct {
	Candidates []int
	Target     int
}

func Test_min(test *testing.T) {
	var cases []CompareCase = []CompareCase{
		CompareCase{[]int{1, 2, 3}, 1},
		CompareCase{[]int{1, 2, -13}, -13},
		CompareCase{[]int{-1, -22, -13}, -22},
		CompareCase{[]int{-11, -2, -13}, -13},
		CompareCase{[]int{0}, 0},
		CompareCase{[]int{0, 0}, 0},
	}

	var result int
	var testcase CompareCase
	for _, testcase = range cases {
		if result = min(testcase.Candidates...); result != testcase.Target {
			test.Errorf(
				"min %v is %d, not %d",
				testcase.Candidates,
				testcase.Target,
				result,
			)
		}
	}
}

func Test_max(test *testing.T) {
	var cases []CompareCase = []CompareCase{
		CompareCase{[]int{1, 2, 3}, 3},
		CompareCase{[]int{1, 2, -13}, 2},
		CompareCase{[]int{-1, -22, -13}, -1},
		CompareCase{[]int{-11, -2, -13}, -2},
		CompareCase{[]int{0}, 0},
		CompareCase{[]int{0, 0}, 0},
	}

	var result int
	var testcase CompareCase
	for _, testcase = range cases {
		if result = max(testcase.Candidates...); result != testcase.Target {
			test.Errorf(
				"max %v is %d, not %d",
				testcase.Candidates,
				testcase.Target,
				result,
			)
		}
	}
}
