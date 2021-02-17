package main

import (
	"math"
)

var (
	KEY_DISTANCE = map[rune]Coordinate{
		'1': Coordinate{0.0, 0.0},
		'2': Coordinate{1.0, 0.0},
		'3': Coordinate{2.0, 0.0},
		'4': Coordinate{3.0, 0.0},
		'5': Coordinate{4.0, 0.0},
		'6': Coordinate{5.0, 0.0},
		'7': Coordinate{6.0, 0.0},
		'8': Coordinate{7.0, 0.0},
		'9': Coordinate{8.0, 0.0},
		'0': Coordinate{9.0, 0.0},
		'q': Coordinate{0.0, 1.0},
		'w': Coordinate{1.0, 1.0},
		'e': Coordinate{2.0, 1.0},
		'r': Coordinate{3.0, 1.0},
		't': Coordinate{4.0, 1.0},
		'y': Coordinate{5.0, 1.0},
		'u': Coordinate{6.0, 1.0},
		'i': Coordinate{7.0, 1.0},
		'o': Coordinate{8.0, 1.0},
		'p': Coordinate{9.0, 1.0},
		'a': Coordinate{0.0, 2.0},
		'z': Coordinate{0.0, 3.0},
		's': Coordinate{1.0, 2.0},
		'x': Coordinate{1.0, 3.0},
		'd': Coordinate{2.0, 2.0},
		'c': Coordinate{2.0, 3.0},
		'f': Coordinate{3.0, 2.0},
		'b': Coordinate{4.0, 3.0},
		'm': Coordinate{5.0, 3.0},
		'g': Coordinate{4.0, 2.0},
		'h': Coordinate{5.0, 2.0},
		'j': Coordinate{6.0, 2.0},
		'k': Coordinate{7.0, 2.0},
		'l': Coordinate{8.0, 2.0},
		'v': Coordinate{3.0, 3.0},
		'n': Coordinate{5.0, 3.0},
	}
)

type Coordinate struct {
	X float64
	Y float64
}

func EuclideanDistance(this, that rune) (distance float64) {
	var distanceX = math.Pow(KEY_DISTANCE[this].X-KEY_DISTANCE[that].X, 2)
	var distanceY = math.Pow(KEY_DISTANCE[this].Y-KEY_DISTANCE[that].Y, 2)
	distance = math.Sqrt(distanceX + distanceY)

	return
}

func TotalEuclideanDistance(this, that string) (total float64) {
	var limit int = min(len(this), len(that))
	for limit != 0 {
		limit--
		total += EuclideanDistance(rune(this[limit]), rune(that[limit]))
	}

	return
}

func MeanEuclideanDistance(this, that string) (mean float64) {
	var lengthAverage float64 = float64(len(this)) + float64(len(that))/2
	return TotalEuclideanDistance(this, that) / lengthAverage
}

func levenshteinMatrix(source, destination string) (matrix [][]int) {
	matrix = make([][]int, len(source)+1)
	var index int
	for index = range matrix {
		matrix[index] = make([]int, len(destination)+1)
	}

	index = 0
	for index != len(source) {
		matrix[index+1][0] = index + 1
		index++
	}

	index = 0
	for index != len(destination) {
		matrix[0][index+1] = index + 1
		index++
	}

	return
}

func LevenshteinDistance(source string, destination string) (distance int) {
	switch {
	case source == destination:
		distance = 0
		return
	case source == "":
		distance = len(destination)
		return
	case destination == "":
		distance = len(source)
		return
	}

	var matrix [][]int = levenshteinMatrix(source, destination)

	var sourceD, destD, cost int
	for sourceD = range matrix[1:] {
		for destD = range matrix[sourceD][1:] {
			if source[sourceD] == destination[destD] {
				cost = 0
			} else {
				cost = 1
			}

			matrix[sourceD+1][destD+1] = min(
				matrix[sourceD][destD+1]+1,
				matrix[sourceD+1][destD]+1,
				matrix[sourceD][destD]+cost,
			)
		}
	}

	distance = matrix[len(source)][len(destination)]
	return
}
