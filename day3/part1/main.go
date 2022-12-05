package main

import (
	"fmt"
	"strings"

	"github.com/hpzzz/aoc22/aocutil"
)

func main() {
	prioritySum := 0

	lines := aocutil.ReadLines("input")
	for _, line := range lines {
		charMap := make(map[byte]bool)
		comp1, comp2 := line[0:len(line)/2], line[len(line)/2:]
		for _, char := range comp1 {
			if strings.ContainsRune(comp2, char) {
				charMap[byte(char)] = true
			}
		}
		for key, _ := range charMap {
			prioritySum += aocutil.Priority(key)
		}
	}
	fmt.Println("Sum of priorities:", prioritySum)

}
