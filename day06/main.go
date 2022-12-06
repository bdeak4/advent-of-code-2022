package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	line := strings.Split(string(file), "\n")[0]

	signallen := 4
	for i := signallen - 1; i < len(line); i++ {
		signal := ""
		for j := signallen - 1; j >= 0; j-- {
			char := string(line[i-j])
			if strings.Contains(signal, char) {
				break
			}
			signal += char
		}
		if len(signal) == signallen {
			fmt.Println("Part 1:", i+1)
			break
		}
	}

	messagelen := 14
	for i := messagelen - 1; i < len(line); i++ {
		message := ""
		for j := messagelen - 1; j >= 0; j-- {
			char := string(line[i-j])
			if strings.Contains(message, char) {
				break
			}
			message += char
		}
		if len(message) == messagelen {
			fmt.Println("Part 2:", i+1)
			break
		}
	}

}
