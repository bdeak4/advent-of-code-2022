package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	ascore := 0
	bscore := 0

	for _, line := range lines {
		strategy := strings.Split(line, " ")
		a := strategy[0]
		b := strategy[1]

		ascore += shapeScore(a)
		bscore += shapeScore(b)

		if isWin(a, b) {
			ascore += 6
			continue
		}
		if isWin(b, a) {
			bscore += 6
			continue
		}
		ascore += 3
		bscore += 3
	}

	fmt.Println("Part 1:", bscore)
}

func isWin(a, b string) bool {
	return (isRock(a) && isScissors(b)) || (isPaper(a) && isRock(b)) || (isScissors(a) && isPaper(b))
}

func isRock(x string) bool {
	return x == "A" || x == "X"
}

func isPaper(x string) bool {
	return x == "B" || x == "Y"
}

func isScissors(x string) bool {
	return x == "C" || x == "Z"
}

func shapeScore(shape string) int {
	if isRock(shape) {
		return 1
	}
	if isPaper(shape) {
		return 2
	}
	if isScissors(shape) {
		return 3
	}
	return 0
}

func part2() {
	file, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	score := 0

	for _, line := range lines {
		strategy := strings.Split(line, " ")
		a := strategy[0]
		outcome := strategy[1]

		// win
		if outcome == "Z" {
			score += shapeScore(winShape(a))
			score += 6
			continue
		}
		// lose
		if outcome == "X" {
			score += shapeScore(loseShape(a))
			continue
		}
		// draw
		score += shapeScore(a)
		score += 3
	}

	fmt.Println("Part 2:", score)

}

func winShape(shape string) string {
	if isRock(shape) {
		return "B"
	}
	if isPaper(shape) {
		return "C"
	}
	if isScissors(shape) {
		return "A"
	}
	return ""
}

func loseShape(shape string) string {
	if isRock(shape) {
		return "C"
	}
	if isPaper(shape) {
		return "A"
	}
	if isScissors(shape) {
		return "B"
	}
	return ""
}
