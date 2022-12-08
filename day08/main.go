package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	sum := 0
	for i, line := range lines {
		for j, tree := range line {
			// edge
			if i == 0 || i == len(lines)-1 || j == 0 || j == len(line)-1 {
				sum += 1
				continue
			}

			// left
			if horizontal_max(line[:j], tree) {
				sum += 1
				continue
			}

			// right
			if horizontal_max(line[j+1:], tree) {
				sum += 1
				continue
			}

			// top
			if vertical_max(lines[:i], j, tree) {
				sum += 1
				continue
			}

			// bottom
			if vertical_max(lines[i+1:], j, tree) {
				sum += 1
				continue
			}
		}
	}
	fmt.Println("Part 1:", sum)

	max_scenic_score := 0
	for i, line := range lines {
		for j, tree := range line {
			// edge
			if i == 0 || i == len(lines)-1 || j == 0 || j == len(line)-1 {
				continue
			}

			left := 0
			for left < len(line[:j]) {
				if byte_to_int(line[j-1-left]) >= rune_to_int(tree) {
					left++
					break
				}
				left++
			}

			right := 0
			for right < len(line[j+1:]) {
				if byte_to_int(line[j+1+right]) >= rune_to_int(tree) {
					right++
					break
				}
				right++
			}

			top := 0
			for top < len(line[:i]) {
				if byte_to_int(lines[i-1-top][j]) >= rune_to_int(tree) {
					top++
					break
				}
				top++
			}

			bottom := 0
			for bottom < len(line[i+1:]) {
				if byte_to_int(lines[i+1+bottom][j]) >= rune_to_int(tree) {
					bottom++
					break
				}
				bottom++
			}

			scenic_score := left * right * top * bottom
			if scenic_score > max_scenic_score {
				max_scenic_score = scenic_score
			}
		}
	}
	fmt.Println("Part 2:", max_scenic_score)
}

func horizontal_max(trees string, tree rune) bool {
	height := rune_to_int(tree)
	for _, t := range trees {
		h := rune_to_int(t)
		if h >= height {
			return false
		}
	}
	return true
}

func vertical_max(lines []string, idx int, tree rune) bool {
	height := rune_to_int(tree)
	for _, line := range lines {
		h := rune_to_int(rune(line[idx]))
		if h >= height {
			return false
		}
	}
	return true
}

func rune_to_int(x rune) int {
	return int(x - '0')
}

func byte_to_int(x byte) int {
	n, _ := strconv.Atoi(string(x))
	return n
}
