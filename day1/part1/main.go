package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	currentElf := 1
	mostCaloriesElf := 1
	currentCalories := 0
	mostCalories := 0
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Compare(text, "") == 0 {
			if currentCalories > mostCalories {
				mostCalories = currentCalories
				mostCaloriesElf = currentElf
			}
			currentCalories = 0
			currentElf += 1
			continue
		} else {
			calories, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			currentCalories += calories
		}
	}
	fmt.Printf("The most calories: %d\nElf number: %d\n", mostCalories, mostCaloriesElf)
}
