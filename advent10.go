package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	// "math"
)

func generateMatrix() [][]string {
	matrix := make([][]string, 6)

	for i := 0; i < 6; i++ {
		row := make([]string, 40)

		for j := 0; j < 40; j++ {
			row[j] = "."
		}

		matrix[i] = row
	}

	return matrix
}

func renderMatrix(m [][]string) string {
	result := ""
	for i := 0; i < len(m); i++ {
		local := ""
		for j := 0; j < len(m[i]); j++ {
			local += m[i][j]
		}

		result += local + "\n"
	}

	return result
}

func two(input string) {
	matrix := generateMatrix()
	X := 1
	currentCycle := 0

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		instruction := parts[0]

		cycles := 1
		add := 0

		if instruction == "addx" {
			cycles = 2
			num, _ := strconv.Atoi(parts[1])
			add = num
		}

		for cycles > 0 {
			i := currentCycle / 40
			j := (currentCycle % 40)
			shouldWrite := j >= X-1 && j <= X+1

			cycles--
			currentCycle++

			if shouldWrite {
				matrix[i][j] = "#"
			}

			if cycles == 0 {
				X += add
			}
		}

	}

	result := renderMatrix(matrix)
	fmt.Println("Second round result")
	fmt.Println(result)
}

func one(input string) {
	period := 0
	signalPeriod := []int{20, 60, 100, 140, 180, 220}
	result := 0
	X := 1
	currentCycle := 0

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		instruction := parts[0]

		cycles := 1
		add := 0

		if instruction == "addx" {
			cycles = 2
			num, _ := strconv.Atoi(parts[1])
			add = num
		}

		for cycles > 0 {
			cycles--
			currentCycle++

			if period < len(signalPeriod) && signalPeriod[period] == currentCycle {
				result += X * currentCycle
				period++
			}

			if cycles == 0 {
				X += add
			}
		}

	}

	fmt.Println("First round result", result)
}

func main() {
	content, err := ioutil.ReadFile("10.txt")

	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	one(input)
	two(input)
}
