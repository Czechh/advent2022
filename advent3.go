package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

func constructAlphaMap() map[int32]int {
	alphaMap := make(map[int32]int)
	priority := 1

	for r := 'a'; r <= 'z'; r++ {
		R := unicode.ToUpper(r)
		alphaMap[r] = priority
		alphaMap[R] = priority + 26

		priority += 1
	}

	return alphaMap
}

func one(input string) {
	alphaMap := constructAlphaMap()
	result := 0

	for _, pair := range strings.Split(input, "\n") {
		charMap := make(map[int32]bool)

		for i, char := range pair {
			if i < (len(pair) / 2) {
				charMap[char] = true
			} else {
				_, present := charMap[char]
				if present {
					fmt.Printf("%c\n", char)
					fmt.Println(pair)
					result += alphaMap[char]
					break
				}
			}
		}
	}

	fmt.Println(result)
}

func getMap(val string) map[int32]bool {
	charMap := make(map[int32]bool)
	for _, c := range val {
		charMap[c] = true
	}

	return charMap
}

func two(input string) {
	alphaMap := constructAlphaMap()
	result := 0
	lines := strings.Split(input, "\n")

	for i := 0; i < len(lines); i += 3 {
		map0 := getMap(lines[i])
		map1 := getMap(lines[i+1])

		fmt.Println(i)

		for _, c := range lines[i+2] {
			_, present0 := map0[c]
			_, present1 := map1[c]

			if present0 && present1 {
				result += alphaMap[c]
				break
			}
		}
	}

	fmt.Println(result)
}

func main() {

	content, err := ioutil.ReadFile("3.txt")

	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	// one(input)
	two(input)
}
