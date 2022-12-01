package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	r1, err := part1()
	if err != nil {
		panic(err)
	}

	r2, err := part2()
	if err != nil {
		panic(err)
	}

	fmt.Println("Part one:", r1)
	fmt.Println("Part two:", r2)
}

func part1() (int, error) {
	input, err := os.ReadFile("./input")
	if err != nil {
		return 0, err
	}

	max := 0
	for _, elf := range strings.Split(string(input), "\n\n") {
		sum := 0
		for _, cal := range strings.Split(elf, "\n") {
			if len(cal) == 0 {
				continue
			}

			c, err := strconv.Atoi(cal)
			if err != nil {
				return 0, err
			}

			sum += c
		}
		if sum > max {
			max = sum
		}
	}

	return max, nil
}

func part2() (int, error) {
	input, err := os.ReadFile("./input")
	if err != nil {
		return 0, err
	}

	sums := []int{}
	for _, elf := range strings.Split(string(input), "\n\n") {
		sum := 0
		for _, cal := range strings.Split(elf, "\n") {
			if len(cal) == 0 {
				continue
			}

			c, err := strconv.Atoi(cal)
			if err != nil {
				return 0, err
			}

			sum += c
		}
		sums = append(sums, sum)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	cal := sums[0] + sums[1] + sums[2]

	return cal, nil
}
