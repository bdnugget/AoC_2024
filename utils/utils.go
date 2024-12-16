package utils

import (
	"os"
	"strings"
)

func InputLines(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}
