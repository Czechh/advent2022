package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func one(input string) {
	moves := map[string]string{
		"A": "R", // 1
		"B": "P", // 2
		"C": "S", // 3
		"X": "R",
		"Y": "P",
		"Z": "S",
	}
	// win: 6, tie: 3, loss: 0
	meScore := 0

	for _, pair := range strings.Split(input, "\n") {
		parts := strings.Split(pair, " ")
		elf := moves[parts[0]]
		me := moves[parts[1]]

		// winning round
		if elf == "R" {
			if me == "R" {
				meScore += 3
			} else if me == "P" {
				meScore += 6
			} else {
				meScore += 0
			}
		} else if elf == "P" {
			if me == "R" {
				meScore += 0
			} else if me == "P" {
				meScore += 3
			} else {
				meScore += 6
			}
		} else if elf == "S" {
			if me == "R" {
				meScore += 6
			} else if me == "P" {
				meScore += 0
			} else {
				meScore += 3
			}
		}

		if me == "R" {
			meScore += 1
		} else if me == "P" {
			meScore += 2
		} else {
			meScore += 3
		}
	}

	fmt.Println(meScore)
}

func two(input string) {
	moves := map[string]string{
		"A": "R", // 1
		"B": "P", // 2
		"C": "S", // 3
	}
	// win: 6, tie: 3, loss: 0
	meScore := 0

	for _, pair := range strings.Split(input, "\n") {
		parts := strings.Split(pair, " ")
		elf := moves[parts[0]]
		outcome := parts[1]

		if elf == "R" {
			if outcome == "X" {
				meScore += 3
			} else if outcome == "Y" {
				meScore += 1
			} else {
				meScore += 2
			}
		} else if elf == "P" {
			if outcome == "X" {
				meScore += 1
			} else if outcome == "Y" {
				meScore += 2
			} else {
				meScore += 3
			}
		} else if elf == "S" {
			if outcome == "X" {
				meScore += 2
			} else if outcome == "Y" {
				meScore += 3
			} else {
				meScore += 1
			}
		}

		if outcome == "X" {
			meScore += 0
		} else if outcome == "Y" {
			meScore += 3
		} else {
			meScore += 6
		}
	}

	fmt.Println(meScore)
}

func main() {

	content, err := ioutil.ReadFile("2.txt")

	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	one(input)
	two(input)
}
