package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/finntolmie/aoc24/day1"
)

func main() {
	path := "inputs/day1.txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	input := string(bytes)
	ans, err := day1.SolveFirstPart(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Solution 1: %d\n", ans)
	ans, err = day1.SolveSecondPart(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Solution 2: %d\n", ans)
}
