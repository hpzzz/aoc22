package main

import (
	"fmt"
	"os"
	"strings"
)

type Shape uint8

const (
	_ Shape = iota
	Rock
	Paper
	Scissors
)

type GameOutcome int

const (
	Loss GameOutcome = 0
	Draw GameOutcome = 3
	Win  GameOutcome = 6
)

func main() {
	var totalScore int

	games := readGames("input")
	for _, game := range games {
		moves := strings.Split(game, " ")
		move1, _ := ParseShape(moves[0])
		move2, _ := ParseShape(moves[1])
		score := int(move2) + int(move2.Score(move1))
		totalScore += score
	}
	fmt.Println(totalScore)
}

func readGames(inputFile string) []string {
	input, _ := os.ReadFile(inputFile)
	games := strings.Split(string(input), "\n")
	return games
}

func ParseShape(input string) (Shape, bool) {
	switch input {
	case "A", "X":
		return Rock, true
	case "B", "Y":
		return Paper, true
	case "C", "Z":
		return Scissors, true
	default:
		return 0, false
	}
}

func (rps Shape) Score(otherShape Shape) GameOutcome {
	switch rps {
	case Rock:
		switch otherShape {
		case Rock:
			return Draw
		case Paper:
			return Loss
		case Scissors:
			return Win
		}
	case Paper:
		switch otherShape {
		case Rock:
			return Win
		case Paper:
			return Draw
		case Scissors:
			return Loss
		}
	case Scissors:
		switch otherShape {
		case Rock:
			return Loss
		case Paper:
			return Win
		case Scissors:
			return Draw
		}
	}
	return Loss
}
