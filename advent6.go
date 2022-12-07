package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func one(input string) {
	m := make(map[int32]int)
	rstr := []rune(input)

	for i := 0; i < len(rstr); i++ {
		c := rstr[i]
		_, present := m[c]

		if present {
			m[c]++
		} else {
			m[c] = 1
		}

		if i > 13 {
			r := rstr[i-14]
			m[r]--

			if m[r] == 0 {
				delete(m, r)
			}
		}

		if len(m) > 13 {
			fmt.Println(i)
			fmt.Println(m)
			break
		}
	}
}

func main() {
	content, err := ioutil.ReadFile("6.txt")

	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	one(input)
	// two(input)
}
