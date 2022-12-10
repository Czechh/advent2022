package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func adjustTail(head []int, tail []int) []int {
	x := head[0] - tail[0]
	y := head[1] - tail[1]

	if x == 2 || x == -2 {
		if y == 1 || y == -1 {
			return []int{tail[0] + (x / 2), head[1]}
		} else if y == 2 || y == -2 {
			return []int{tail[0] + (x / 2), tail[1] + (y / 2)}
		}

		return []int{tail[0] + (x / 2), tail[1]}
	} else if y == 2 || y == -2 {
		if x == 1 || x == -1 {
			return []int{head[0], tail[1] + (y / 2)}
		} else if x == 2 || x == -2 {
			return []int{tail[0] + (x / 2), tail[1] + (y / 2)}
		}

		return []int{tail[0], tail[1] + (y / 2)}
	}

	return tail
}

func one(input string) {
	head := []int{0, 0}
	tail := []int{0, 0}
	result := make(map[string]int)

	result["0-0"] = 1

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		direction := parts[0]
		num, _ := strconv.Atoi(parts[1])

		for num > 0 {
			if direction == "U" {
				head[1] += 1
			} else if direction == "D" {
				head[1] -= 1
			} else if direction == "L" {
				head[0] -= 1
			} else if direction == "R" {
				head[0] += 1
			}

			tail = adjustTail(head, tail)
			strCoord := strconv.Itoa(tail[0]) + "-" + strconv.Itoa(tail[1])
			result[strCoord] = 1

			num--
		}
	}

	fmt.Println("First Round results", len(result))
}

func two(input string) {
	rope := make([][]int, 10)

	for i := 0; i < 10; i++ {
		rope[i] = []int{0, 0}
	}

	result := make(map[string]int)

	result["0-0"] = 1

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		direction := parts[0]
		num, _ := strconv.Atoi(parts[1])

		for num > 0 {
			if direction == "U" {
				rope[0][1] += 1
			} else if direction == "D" {
				rope[0][1] -= 1
			} else if direction == "L" {
				rope[0][0] -= 1
			} else if direction == "R" {
				rope[0][0] += 1
			}

			for i := 1; i < len(rope); i++ {
				rope[i] = adjustTail(rope[i-1], rope[i])
			}

			strCoord := strconv.Itoa(rope[9][0]) + "-" + strconv.Itoa(rope[9][1])
			result[strCoord] = 1

			num--
		}
	}

	fmt.Println("Second round results", len(result))
}

func main() {
	content, err := ioutil.ReadFile("9.txt")

	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	one(input)
	two(input)
}
