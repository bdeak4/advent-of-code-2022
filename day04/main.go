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

	countContain := 0
	countOverlap := 0

	for _, line := range lines {
		ab := strings.Split(line, ",")
		a := strings.Split(ab[0], "-")
		b := strings.Split(ab[1], "-")

		amin, _ := strconv.Atoi(a[0])
		amax, _ := strconv.Atoi(a[1])
		bmin, _ := strconv.Atoi(b[0])
		bmax, _ := strconv.Atoi(b[1])

		if (amin >= bmin && amax <= bmax) || (amin <= bmin && amax >= bmax) {
			countContain++
		}

		if (amin >= bmin && amin <= bmax) || (bmin >= amin && bmin <= amax) {
			countOverlap++
		}
	}

	fmt.Println("Part 1:", countContain)
	fmt.Println("Part 2:", countOverlap)
}
