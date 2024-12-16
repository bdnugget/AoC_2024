package main

import (
	"fmt"
	"time"

	"github.com/bdnugget/AoC_2024/day01"
	"github.com/bdnugget/AoC_2024/day02"
)

func measureAndPrint[T any](label string, fn func() T) {
	startTime := time.Now()
	res := fn()
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Printf("%s: %-15v\t(%v)\n", label, res, elapsed)
}

func main() {
	measureAndPrint("Day 01 part 01", day01.Part01)
	measureAndPrint("Day 01 part 02", day01.Part02)
	measureAndPrint("Day 02 part 01", day02.Part01)
	measureAndPrint("Day 02 part 02", day02.Part02)
}
