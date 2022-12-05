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
		chunks := strings.Split(game, " ")
		move1, _ := ParseShape(chunks[0])
		outcome, _ := ParseOutcome(chunks[1])
		move2 := move1.expectedMove(outcome)
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

func ParseOutcome(input string) (GameOutcome, bool) {
	switch input {
	case "X":
		return Loss, true
	case "Y":
		return Draw, true
	case "Z":
		return Win, true
	default:
		return 0, false
	}
}

func (rps Shape) expectedMove(expectedOutcome GameOutcome) Shape {
	switch rps {
	case Rock:
		switch expectedOutcome {
		case Loss:
			return Scissors
		case Draw:
			return Rock
		case Win:
			return Paper
		}
	case Paper:
		switch expectedOutcome {
		case Loss:
			return Rock
		case Draw:
			return Paper
		case Win:
			return Scissors
		}
	case Scissors:
		switch expectedOutcome {
		case Loss:
			return Paper
		case Draw:
			return Scissors
		case Win:
			return Rock
		}
	}
	return 0
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
