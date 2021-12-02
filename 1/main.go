package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/tobiaskohlbau/adventofcode2021/pkg/input"
)

func main() {
	if err := execute(); err != nil {
		log.Fatal(err)
	}
}

func execute() error {
	sample := flag.Bool("sample", false, "use sample input")
	flag.Parse()

	filename := "challenge.input"
	if *sample {
		filename = "sample.input"
	}

	lines, err := input.Read(filename)
	if err != nil {
		return fmt.Errorf("error reading input file: %w", err)
	}

	items := []int{}
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			return fmt.Errorf("error parsing line: %w", err)
		}
		items = append(items, i)
	}

	// part one
	count := 0
	for i, item := range items[1:] {
		if item > items[i] {
			count++
		}
	}

	fmt.Println(count)

	// part two
	count = 0
	lastSum := -1
	for i := 0; i < len(items)-2; i++ {
		sum := items[i+0] + items[i+1] + items[i+2]
		if lastSum != -1 && sum > lastSum {
			count++
		}
		lastSum = sum
	}

	fmt.Println(count)

	return nil
}
