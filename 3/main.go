package main

import (
	"flag"
	"fmt"
	"log"
	"math"

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

	width := len(lines[0])
	height := len(lines)

	grid := make([]int, width*height)

	for y, line := range lines {
		for x, c := range line {
			pos := y*width + x
			if c == '1' {
				grid[pos] = 1
			} else if c == '0' {
				grid[pos] = 0
			}
		}
	}

	// part one
	gammarate := []int{}
	epsilonrate := []int{}
	searchrange := []int{}
	for i := 0; i < height; i++ {
		searchrange = append(searchrange, i)
	}
	for x := 0; x < width; x++ {
		zeroes, ones := countValues(grid, searchrange, width, x)
		if zeroes > ones {
			gammarate = append(gammarate, 0)
			epsilonrate = append(epsilonrate, 1)
		}
		if ones > zeroes {
			gammarate = append(gammarate, 1)
			epsilonrate = append(epsilonrate, 0)
		}
	}

	fmt.Println(convertToDecimal(gammarate) * convertToDecimal(epsilonrate))

	// part two
	checkOxygen := func(zeroes, ones, value int) bool {
		if zeroes > ones && value == 0 {
			return true
		}
		if ones > zeroes && value == 1 {
			return true
		}
		if ones == zeroes && value == 1 {
			return true
		}
		return false
	}
	oxygenRating := searchValue(grid, height, width, checkOxygen)
	co2Rating := searchValue(grid, height, width, func(zeroes, ones, value int) bool {
		return !checkOxygen(zeroes, ones, value)
	})

	fmt.Println(convertToDecimal(oxygenRating) * convertToDecimal(co2Rating))

	return nil
}

func countValues(grid []int, searchrange []int, width, x int) (int, int) {
	zeroes := 0
	ones := 0
	for _, y := range searchrange {
		value := grid[y*width+x]
		if value == 0 {
			zeroes++
		}
		if value == 1 {
			ones++
		}
	}

	return zeroes, ones
}

func searchValue(grid []int, height, width int, fn func(zeroes, ones, value int) bool) []int {
	searchrange := []int{}
	for i := 0; i < height; i++ {
		searchrange = append(searchrange, i)
	}
	x := 0
	row := -1
	for {
		if len(searchrange) == 1 {
			row = searchrange[0]
			break
		}

		zeroes, ones := countValues(grid, searchrange, width, x)

		tmpSearchrange := []int{}
		for _, y := range searchrange {
			value := grid[y*width+x]
			if fn(zeroes, ones, value) {
				tmpSearchrange = append(tmpSearchrange, y)
			}
		}

		searchrange = tmpSearchrange
		x++
	}

	if row != -1 {
		return grid[row*width : row*width+width]
	}
	return []int{}
}

func convertToDecimal(binary []int) int {
	decimal := 0
	highest := len(binary) - 1
	for i, b := range binary {
		decimal += b * int(math.Pow(2, float64(highest-i)))
	}
	return decimal
}
