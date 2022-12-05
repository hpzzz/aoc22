package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	caloryList := []int{}
	currentCalories := 0
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Compare(text, "") == 0 {
			caloryList = append(caloryList, currentCalories)
			currentCalories = 0
			continue
		} else {
			calories, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			currentCalories += calories
		}
	}
	sort.Slice(caloryList, func(i, j int) bool {
		return caloryList[i] > caloryList[j]
	})
	topCalorySum := caloryList[0] + caloryList[1] + caloryList[2]
	fmt.Println(topCalorySum)
}
