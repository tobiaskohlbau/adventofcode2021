package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

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

	type Movement struct {
		Direction string
		Count     int
	}

	movements := make([]Movement, 0)
	for _, line := range lines {
		items := strings.Split(line, " ")
		if len(items) != 2 {
			return fmt.Errorf("invalid line: %s", line)
		}
		count, err := strconv.Atoi(items[1])
		if err != nil {
			return fmt.Errorf("error parsing movement count: %w", err)
		}
		movements = append(movements, Movement{Direction: items[0], Count: count})
	}

	// part one
	positionHorizontal := 0
	depth := 0

	for _, movement := range movements {
		if movement.Direction == "forward" {
			positionHorizontal += movement.Count
		}
		if movement.Direction == "down" {
			depth += movement.Count
		}
		if movement.Direction == "up" {
			depth -= movement.Count
		}
	}

	fmt.Printf("Horizontal position: %d\n", positionHorizontal)
	fmt.Printf("Depth: %d\n", depth)
	fmt.Printf("Total: %d\n", positionHorizontal*depth)

	// part two
	positionHorizontal = 0
	depth = 0
	aim := 0
	for _, movement := range movements {
		if movement.Direction == "forward" {
			positionHorizontal += movement.Count
			depth += aim * movement.Count
		}
		if movement.Direction == "down" {
			aim += movement.Count
		}
		if movement.Direction == "up" {
			aim -= movement.Count
		}
	}
	fmt.Printf("Horizontal position: %d\n", positionHorizontal)
	fmt.Printf("Depth: %d\n", depth)
	fmt.Printf("Total: %d\n", positionHorizontal*depth)

	return nil
}
