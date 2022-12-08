package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func one(input string) {
	rows := strings.Fields(input)

	matrix := make([][]int, len(rows))

	for i, line := range strings.Split(input, "\n") {
		chars := []rune(line)
		row := make([]int, len(rows))
		for j := 0; j < len(chars); j++ {
			num, _ := strconv.Atoi(string(chars[j]))
			row[j] = num
		}
		matrix[i] = row
	}

	result := 0

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows); j++ {
			current := matrix[i][j]
			if i == 0 || i == len(rows)-1 || j == 0 || j == len(rows)-1 {
				result++
			} else {
				shouldAdd := true

				// This is bad lol
				for n := i + 1; n < len(rows); n++ {
					if matrix[n][j] >= current {
						shouldAdd = false
						break
					}
				}

				if shouldAdd {
					result++
					continue
				}

				shouldAdd = true

				for n := i - 1; n >= 0; n-- {
					if matrix[n][j] >= current {
						shouldAdd = false
						break
					}
				}

				if shouldAdd {
					result++
					continue
				}

				shouldAdd = true

				for n := j + 1; n < len(rows); n++ {
					if matrix[i][n] >= current {
						shouldAdd = false
						break
					}
				}

				if shouldAdd {
					result++
					continue
				}

				shouldAdd = true

				for n := j - 1; n >= 0; n-- {
					if matrix[i][n] >= current {
						shouldAdd = false
						break
					}
				}

				if shouldAdd {
					result++
					continue
				}
			}
		}
	}

	fmt.Println("First section visible trees", result)
}

func two(input string) {
	rows := strings.Fields(input)
	matrix := make([][]int, len(rows))
	max := 0

	for i, line := range strings.Split(input, "\n") {
		chars := []rune(line)
		row := make([]int, len(rows))
		for j := 0; j < len(chars); j++ {
			num, _ := strconv.Atoi(string(chars[j]))
			row[j] = num
		}
		matrix[i] = row
	}

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows); j++ {
			current := matrix[i][j]
			scores := make([]int, 4) // [T, B, R, L]

			for n := i + 1; n < len(rows); n++ {
				scores[0]++
				if matrix[n][j] >= current {
					break
				}
			}

			for n := i - 1; n >= 0; n-- {
				scores[1]++
				if matrix[n][j] >= current {
					break
				}
			}

			for n := j + 1; n < len(rows); n++ {
				scores[2]++
				if matrix[i][n] >= current {
					break
				}
			}

			for n := j - 1; n >= 0; n-- {
				scores[3]++
				if matrix[i][n] >= current {
					break
				}
			}

			localScore := scores[0] * scores[1] * scores[2] * scores[3]

			if localScore > max {
				max = localScore
			}
		}
	}

	fmt.Println("Second section score", max)
}

func main() {
	content, err := ioutil.ReadFile("8.txt")

	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	one(input)
	two(input)
}
