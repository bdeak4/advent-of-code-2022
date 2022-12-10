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

	cycle := 0
	x := 1
	sum := 0
	for _, line := range lines {
		linesplit := strings.Split(line, " ")
		switch linesplit[0] {
		case "noop":
			cycle++
			sum += measure_signal(cycle, x)
		case "addx":
			cycle++
			sum += measure_signal(cycle, x)

			cycle++
			sum += measure_signal(cycle, x)
			increment, _ := strconv.Atoi(linesplit[1])
			x += increment
		}

	}

	fmt.Println("Part 1:", sum)

	pos := 0
	sprite := 0
	var crt [200]string
	for i := range crt {
		crt[i] = "."
	}
	for _, line := range lines {
		linesplit := strings.Split(line, " ")
		switch linesplit[0] {
		case "noop":
			if sprite <= 3 {
				crt[pos+sprite] = "#"
				sprite++
			}
			//pos++

		case "addx":
			if sprite <= 3 {
				crt[pos+sprite] = "#"
				sprite++
			}
			//pos++

			increment, _ := strconv.Atoi(linesplit[1])
			pos += increment
			if pos < 0 {
				pos = 0
			}
			sprite = 0
			//for ; increment > 0; increment-- {
			//	crt[pos] = "."
			//	pos++
			//}
		}

	}

	for i := 0; i <= pos; i += 40 {
		fmt.Println(crt[i : i+40])
	}
}

func contains(s []int, n int) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}
	return false
}

func measure_signal(cycle, x int) int {
	measure_cycles := []int{20, 60, 100, 140, 180, 220}

	if !contains(measure_cycles, cycle) {
		return 0
	}

	return cycle * x
}
