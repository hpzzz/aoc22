package aocoutil

import (
	"os"
	"strings"
)

func readLines(inputFile string) []string {
	input, _ := os.ReadFile(inputFile)
	lines := strings.Split(string(input), "\n")
	return lines
}
