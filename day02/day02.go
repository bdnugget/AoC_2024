package day02

import (
	_ "fmt"
	"slices"
	"strings"

	"github.com/bdnugget/AoC_2024/utils"
)

func Part01() int {
	reportLines := utils.InputLines("./day02/input")
	var reports [][]int
	for _, report := range reportLines {
		levels := strings.Fields(report)
		var row []int
		for i, level := range levels {
			levelInt := utils.MustAtoi(level)
			if i == 0 {
				row = append(row, levelInt)
			} else if i > 0 && utils.AbsInt(levelInt-row[i-1]) >= 1 && utils.AbsInt(levelInt-row[i-1]) <= 3 {
				row = append(row, levelInt)
			} else {
				row = nil
				break
			}
		}
		if row == nil {
			continue
		}
		if slices.IsSorted(row) || slices.IsSortedFunc(row, func(a, b int) int { return b - a }) {
			reports = append(reports, row)
		}

	}
	return len(reports)
}

func Part02() int {
	return 69
}
