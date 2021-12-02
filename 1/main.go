package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	partOne()
	partTwo(3)
}

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	previous := math.MaxUint32
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		if current > previous {
			total++
		}
		previous = current
	}
	fmt.Printf("Solution to part one: %d\n", total)
}

func partTwo(window int) {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	queue := make([]int, window)
	total := -window
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		previous := queue[0]
		queue = queue[1:]
		if current > previous {
			total++
		}
		queue = append(queue, current)
	}
	fmt.Printf("Solution to part two: %d\n", total)
}
