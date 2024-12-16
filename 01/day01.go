package day01

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	fmt.Println(InputLines("input"))
}

func InputLines(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}

func Part01() int {
	return 69
}
