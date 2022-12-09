package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	X, Y int
}

func main() {
	file, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	h := Position{0, 0}
	t := Position{0, 0}
	tcount := make(map[Position]int)
	for _, line := range lines {
		linesplit := strings.Split(line, " ")
		direction := linesplit[0]
		steps, _ := strconv.Atoi(linesplit[1])

		for ; steps > 0; steps-- {
			movePosition(&h, direction)
			followPosition(&t, &h, direction, tcount)
		}
	}

	sum := 0
	for _ = range tcount {
		sum++
	}
	fmt.Println("Part 1:", sum)

	var positions [10]Position
	tcount = make(map[Position]int)
	for _, line := range lines {
		linesplit := strings.Split(line, " ")
		direction := linesplit[0]
		steps, _ := strconv.Atoi(linesplit[1])

		for ; steps > 0; steps-- {
			for i := range positions {
				if i == 0 {
					movePosition(&positions[0], direction)
					continue
				}

				if i == len(positions)-1 {
					followPosition(&positions[i], &positions[i-1], direction, tcount)
					continue
				}

				followPosition(&positions[i], &positions[i-1], direction, nil)
			}
		}
	}

	sum = 0
	for _ = range tcount {
		sum++
	}
	fmt.Println("Part 2:", sum)
}

func followPosition(p1, p2 *Position, direction string, count map[Position]int) {
	if isConnected(*p1, *p2) {
		return
	}

	// on same vertical/horizontal line, it is enough to move tail in same direction
	if p1.X == p2.X || p1.Y == p2.Y {
		movePosition(p1, direction)
		if count != nil {
			count[*p1]++
		}
		return
	}

	if direction == "U" || direction == "D" {
		p1.X = p2.X
	}

	if direction == "R" || direction == "L" {
		p1.Y = p2.Y
	}

	movePosition(p1, direction)
	if count != nil {
		count[*p1]++
	}
}

func movePosition(p *Position, direction string) {
	switch direction {
	case "R":
		p.X++
	case "L":
		p.X--
	case "U":
		p.Y++
	case "D":
		p.Y--
	}
}

func isConnected(p1, p2 Position) bool {
	return p1.X+1 >= p2.X && p1.X-1 <= p2.X && p1.Y+1 >= p2.Y && p1.Y-1 <= p2.Y
}
