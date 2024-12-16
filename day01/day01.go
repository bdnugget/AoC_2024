package day01

import (
	"math"
	"slices"
	"strings"

	"github.com/bdnugget/AoC_2024/utils"
)

func getSortedCols(lines []string) (leftCol []int, rightCol []int) {
	for _, line := range lines {
		left, right, found := strings.Cut(line, "   ")
		if found {
			leftCol = append(leftCol, utils.MustAtoi(left))
			rightCol = append(rightCol, utils.MustAtoi(right))
		}
	}
	slices.Sort(leftCol)
	slices.Sort(rightCol)

	return leftCol, rightCol
}

func Part01() int {
	lines := utils.InputLines("./day01/input")
	leftCol, rightCol := getSortedCols(lines)

	var diffCols []int
	for i, val := range leftCol {
		diffCols = append(diffCols, int(math.Abs(float64((val - rightCol[i])))))
	}

	sumDiff := 0
	for _, diff := range diffCols {
		sumDiff += diff
	}

	return sumDiff
}

func Part02() int {
	lines := utils.InputLines("./day01/input")
	leftCol, rightCol := getSortedCols(lines)

	freqsRightCol := make(map[int]int)
	for _, num := range rightCol {
		freqsRightCol[num] += 1
	}

	var similarity int
	for _, num := range leftCol {
		similarity += freqsRightCol[num] * num
	}

	return similarity
}
