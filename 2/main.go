package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
}

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	forwardness := 0
	depth := 0
	for scanner.Scan() {
		direction, magnitude := parseCommand(scanner.Text())
		switch direction {
		case "forward":
			forwardness = goForward(forwardness, magnitude)
		case "backward":
			forwardness = goForward(forwardness, -magnitude)
		case "up":
			depth = goDown(depth, -magnitude)
		case "down":
			depth = goDown(depth, magnitude)
		default:
			panic("Unexpected input")
		}
	}
	fmt.Printf("Forwardness: %d\n", forwardness)
	fmt.Printf("Depth: %d\n", depth)
}

func parseCommand(input string) (string, int) {
	slice := strings.Split(input, " ")
	if len(slice) != 2 {
		panic("Unexpected input")
	}
	parsed, err := strconv.Atoi(slice[1])
	if err != nil {
		panic("Unexpected input")
	}
	return slice[0], parsed
}

func goForward(forwardness, magnitude int) int {
	return forwardness + magnitude
}

func goDown(depth, magnitude int) int {
	return max(depth+magnitude, 0)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
