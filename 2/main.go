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
	partTwo()
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

func partTwo() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	forwardness := 0
	depth := 0
	aim := 0
	for scanner.Scan() {
		direction, magnitude := parseCommand(scanner.Text())
		switch direction {
		case "forward":
			forwardness, depth = goForwardWithAim(forwardness, depth, aim, magnitude)
		case "backward":
			forwardness, depth = goForwardWithAim(forwardness, depth, aim, -magnitude)
		case "up":
			aim = adjustAim(aim, -magnitude)
		case "down":
			aim = adjustAim(aim, magnitude)
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

func goForwardWithAim(forwardness, depth, aim, magnitude int) (int, int) {
	depth = max(depth+aim*magnitude, 0)
	forwardness = forwardness + magnitude
	return forwardness, depth
}

func adjustAim(aim, magnitude int) int {
	return aim + magnitude
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
