package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
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
	fmt.Printf("%d\n", total)
}
