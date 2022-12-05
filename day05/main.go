package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	stackSize  = 8
	stackCount = 9
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
	lines := strings.Split(string(file), "\n")

	stacks := make([][]string, stackCount)
	for i := stackSize - 1; i >= 0; i-- {
		stacki := 0
		for j := 1; j < 4*stackCount; j += 4 {
			if unicode.IsLetter(rune(lines[i][j])) {
				stacks[stacki] = append(stacks[stacki], string(lines[i][j]))
			}
			stacki++

		}
	}

	for _, instruction := range lines[stackCount+1 : len(lines)-1] {
		instructionParts := strings.Split(instruction, " ")
		count, _ := strconv.Atoi(instructionParts[1])
		srcnum, _ := strconv.Atoi(instructionParts[3])
		dstnum, _ := strconv.Atoi(instructionParts[5])

		srci := srcnum - 1
		dsti := dstnum - 1

		for i := 0; i < count; i++ {
			if len(stacks[srci]) == 0 {
				break
			}
			stacks[dsti] = append(stacks[dsti], stacks[srci][len(stacks[srci])-1])
			stacks[srci] = stacks[srci][:len(stacks[srci])-1]
		}
	}

	result := ""
	for _, stack := range stacks {
		if len(stack) > 0 {
			result += stack[len(stack)-1]
		}
	}
	fmt.Println("Part 1:", result)
}

func part2() {
	file, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")

	stacks := make([][]string, stackCount)
	for i := stackSize - 1; i >= 0; i-- {
		stacki := 0
		for j := 1; j < 4*stackCount; j += 4 {
			if unicode.IsLetter(rune(lines[i][j])) {
				stacks[stacki] = append(stacks[stacki], string(lines[i][j]))
			}
			stacki++

		}
	}

	for _, instruction := range lines[stackCount+1 : len(lines)-1] {
		instructionParts := strings.Split(instruction, " ")
		count, _ := strconv.Atoi(instructionParts[1])
		srcnum, _ := strconv.Atoi(instructionParts[3])
		dstnum, _ := strconv.Atoi(instructionParts[5])

		srci := srcnum - 1
		dsti := dstnum - 1

		if count > len(stacks[srci]) {
			count = len(stacks[srci])
		}

		stacks[dsti] = append(stacks[dsti], stacks[srci][len(stacks[srci])-count:]...)
		stacks[srci] = stacks[srci][:len(stacks[srci])-count]
	}

	result := ""
	for _, stack := range stacks {
		if len(stack) > 0 {
			result += stack[len(stack)-1]
		}
	}
	fmt.Println("Part 2:", result)
}
