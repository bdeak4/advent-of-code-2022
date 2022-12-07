package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	DISK_SIZE   = 70_000_000
	UPDATE_SIZE = 30_000_000
)

func main() {
	file, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	fs := make(map[string]int)
	pwd := []string{}
	for _, line := range lines {
		args := strings.Split(line, " ")
		if args[1] == "cd" {
			dir := args[2]
			if dir == ".." {
				pwd = pwd[:len(pwd)-1]
				continue
			}
			pwd = append(pwd, dir)
			fs[strings.Join(pwd, "/")] = 0
			continue
		}
		if size, err := strconv.Atoi(args[0]); err == nil {
			for i, _ := range pwd {
				fs[strings.Join(pwd[:i+1], "/")] += size
			}
		}
	}

	sum := 0
	for _, size := range fs {
		if size < 100000 {
			sum += size
		}
	}
	fmt.Println("Part 1:", sum)

	unused_space := DISK_SIZE - fs["/"]
	free_required := UPDATE_SIZE - unused_space
	dir_size := fs["/"]
	for _, size := range fs {
		if size > free_required && size < dir_size {
			dir_size = size
		}
	}
	fmt.Println("Part 2:", dir_size)
}
