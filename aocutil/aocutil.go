package aocutil

import (
	"os"
	"strings"
)

func ReadLines(inputFile string) []string {
	input, _ := os.ReadFile(inputFile)
	lines := strings.Split(string(input), "\n")
	return lines
}

func Priority(b byte) int {
	if 'a' <= b && b <= 'z' {
		return int(b-'a') + 1
	} else if 'A' <= b && b <= 'Z' {
		return int(b-'A') + 27
	} else {
		return 0
	}
}
