package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

type Monkey struct {
	items     []int
	operation func(int) int
	test      func(int) int
	divisor   int
}

func getInitialState() []Monkey {
	monkey0 := Monkey{
		items:     []int{63, 57},
		operation: func(old int) int { return old * 11 },
		test: func(val int) int {
			if val%7 == 0 {
				return 6
			}
			return 2
		},
		divisor: 7,
	}

	monkey1 := Monkey{
		items:     []int{82, 66, 87, 78, 77, 92, 83},
		operation: func(old int) int { return old + 1 },
		test: func(val int) int {
			if val%11 == 0 {
				return 5
			}
			return 0
		},
		divisor: 11,
	}

	monkey2 := Monkey{
		items:     []int{97, 53, 53, 85, 58, 54},
		operation: func(old int) int { return old * 7 },
		test: func(val int) int {
			if val%13 == 0 {
				return 4
			}
			return 3
		},
		divisor: 13,
	}

	monkey3 := Monkey{
		items:     []int{50},
		operation: func(old int) int { return old + 3 },
		test: func(val int) int {
			if val%3 == 0 {
				return 1
			}
			return 7
		},
		divisor: 3,
	}

	monkey4 := Monkey{
		items:     []int{64, 69, 52, 65, 73},
		operation: func(old int) int { return old + 6 },
		test: func(val int) int {
			if val%17 == 0 {
				return 3
			}
			return 7
		},
		divisor: 17,
	}

	monkey5 := Monkey{
		items:     []int{57, 91, 65},
		operation: func(old int) int { return old + 5 },
		test: func(val int) int {
			if val%2 == 0 {
				return 0
			}
			return 6
		},
		divisor: 2,
	}

	monkey6 := Monkey{
		items:     []int{67, 91, 84, 78, 60, 69, 99, 83},
		operation: func(old int) int { return old * old },
		test: func(val int) int {
			if val%5 == 0 {
				return 2
			}
			return 4
		},
		divisor: 5,
	}

	monkey7 := Monkey{
		items:     []int{58, 78, 69, 65},
		operation: func(old int) int { return old + 7 },
		test: func(val int) int {
			if val%19 == 0 {
				return 5
			}
			return 1
		},
		divisor: 19,
	}

	result := []Monkey{monkey0, monkey1, monkey2, monkey3, monkey4, monkey5, monkey6, monkey7}

	return result
}

func one(input string) {
	state := getInitialState()
	results := []int{0, 0, 0, 0, 0, 0, 0, 0}

	round := 0

	for round < 20 {
		for i, monkey := range state {
			for _, item := range monkey.items {
				newVal := monkey.operation(item) / 3
				monkeyTarget := monkey.test(newVal)
				state[monkeyTarget].items = append(state[monkeyTarget].items, newVal)
				results[i]++
			}

			state[i].items = make([]int, 0)
		}

		round++
	}

	sort.Ints(results)
	shenanigans := results[len(results)-2] * results[len(results)-1]

	fmt.Println("First round result", results, shenanigans)
}

func two(input string) {
	state := getInitialState()
	results := []int{0, 0, 0, 0, 0, 0, 0, 0}

	mod := 1
	for i := range state {
		mod *= state[i].divisor
	}

	round := 0

	for round < 10000 {
		for i, monkey := range state {
			for _, item := range monkey.items {
				newVal := monkey.operation(item) % mod
				monkeyTarget := monkey.test(newVal)
				state[monkeyTarget].items = append(state[monkeyTarget].items, newVal)
				results[i]++
			}

			state[i].items = make([]int, 0)
		}

		round++
	}

	sort.Ints(results)
	shenanigans := results[len(results)-2] * results[len(results)-1]

	fmt.Println("Second round result", results, shenanigans)
}

func main() {
	content, err := ioutil.ReadFile("11.txt")

	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	one(input)
	two(input)
}
