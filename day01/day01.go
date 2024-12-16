package day01

import (
	"os"
	"strings"
	"strconv"
	//"fmt"
	"slices"
	"math"
)

func InputLines(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}

func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

func Part01() int {
	lines := InputLines("./day01/input")
	var leftCol, rightCol []int
	for _, line := range lines {
		left, right, found := strings.Cut(line, "   ")
		if found {
			leftCol = append(leftCol, atoi(left))
			rightCol = append(rightCol, atoi(right))
		}
	}

	slices.Sort(leftCol)
	slices.Sort(rightCol)

	var diffCols []int
	for i, val := range leftCol {
		diffCols = append(diffCols, int(math.Abs(float64((val - rightCol[i])))))
	}

	sumDiff := 0
	for _, diff := range diffCols {
		sumDiff += diff
	}

	//fmt.Printf("%v", sumDiff)
	return sumDiff
}
