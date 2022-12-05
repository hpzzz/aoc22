package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hpzzz/aoc22/aocutil"
)

func main() {
	prioritySum := 0

	groupRucksacks := []string{}

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		if len(str) == 0 {
			continue
		}
		groupRucksacks = append(groupRucksacks, str)
		if len(groupRucksacks) == groupSize {
			comp1, comp2, comp3 := groupRucksacks[0], groupRucksacks[1], groupRucksacks[2]
			fmt.Println(comp1, comp2, comp3)
			for _, char := range comp1 {
				if strings.ContainsRune(comp2, char) && strings.ContainsRune(comp3, char) {
					prioritySum += aocutil.Priority(byte(char))
					break
				}
			}
			groupRucksacks = groupRucksacks[:0]
		}
	}
	fmt.Println(prioritySum)
}
