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

	sum := 0
	for _, rucksack := range lines {
		firstCompartment := rucksack[0 : len(rucksack)/2]
		secondCompartment := rucksack[len(rucksack)/2:]
		sum += strings.IndexRune(alphabet, findChar(firstCompartment, secondCompartment))
	}

	fmt.Println("Part 1:", sum)
}

const alphabet = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func findChar(str1, str2 string) rune {
	for _, char := range str1 {
		if strings.IndexRune(str2, char) > -1 {
			return char
		}
	}
	return ' '
}

func part2() {
	file, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	sum := 0
	for i, _ := range lines {
		if i%3 != 0 {
			continue
		}
		sum += strings.IndexRune(alphabet, findChar3(lines[i], lines[i+1], lines[i+2]))
	}

	fmt.Println("Part 2:", sum)
}

func findChar3(str1, str2, str3 string) rune {
	for _, char := range str1 {
		if strings.IndexRune(str2, char) > -1 && strings.IndexRune(str3, char) > -1 {
			return char
		}
	}
	return ' '
}
