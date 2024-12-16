package day01

import (
	"fmt"
	"testing"
)

func TestInputLines(t *testing.T) {
	got := InputLines("testinput")
	want := []string{
		"aaa",
		"bbb",
		"ccc",
		"",
	}

	for i, str := range got {
		fmt.Println(i)
		if str != want[i] {
			t.Errorf("got %q, wanted %q", str, want[i])
		}
	}
}
