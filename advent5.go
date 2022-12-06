package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
    "strconv"
)

//     [H]         [D]     [P]        
// [W] [B]         [C] [Z] [D]        
// [T] [J]     [T] [J] [D] [J]        
// [H] [Z]     [H] [H] [W] [S]     [M]
// [P] [F] [R] [P] [Z] [F] [W]     [F]
// [J] [V] [T] [N] [F] [G] [Z] [S] [S]
// [C] [R] [P] [S] [V] [M] [V] [D] [Z]
// [F] [G] [H] [Z] [N] [P] [M] [N] [D]
//  1   2   3   4   5   6   7   8   9 


func getInitialState() [][]string {
    results := [][]string{
        {"W", "T", "H", "P", "J", "C", "F"},
        {"H","B","J","Z","F","V","R","G"},
        {"R", "T", "P", "H"},
        {"T", "H", "P", "N", "S", "Z"},
        {"D", "C", "J", "H", "Z", "F", "V", "N"},
        {"Z", "D", "W", "F", "G", "M", "P"},
        {"P", "D", "J", "S", "W", "Z", "V", "M"},
        {"S", "D", "N"},
        {"M", "F", "S", "Z", "D"},
    }

    for _, line := range results {
        for i, j := 0, len(line)-1; i < j; i, j = i+1, j-1 {
            line[i], line[j] = line[j], line[i]
        }
    }

    return results
}

func one(input string) {
    state := getInitialState();

    for _, line := range strings.Split(input, "\n") {
        parts := strings.Split(line, " ")

        counts, _ := strconv.Atoi(parts[1])
        origin, _ := strconv.Atoi(parts[3])
        target, _ := strconv.Atoi(parts[5])

        for i := 0; i < counts; i++ {
            n := len(state[origin - 1]) - 1
            temp := state[origin - 1][n]
            state[origin - 1] = state[origin - 1][:n]
            state[target - 1] = append(state[target - 1], temp)
        }
    }

    result := ""

    for _, line := range state {
        result += line[len(line) - 1]
    }

    fmt.Println(result)
}

func two(input string) {
    state := getInitialState();

    for _, line := range strings.Split(input, "\n") {
        parts := strings.Split(line, " ")

        counts, _ := strconv.Atoi(parts[1])
        origin, _ := strconv.Atoi(parts[3])
        target, _ := strconv.Atoi(parts[5])

        n := len(state[origin - 1])
        toMove := state[origin - 1][n - counts:]
        state[origin - 1] = state[origin - 1][:n - counts]

        for _, p := range toMove {
            state[target - 1] = append(state[target - 1], p)
        }
    }

    result := ""

    for _, line := range state {
        result += line[len(line) - 1]
    }

    fmt.Println(result)
}

func main() {
    content, err := ioutil.ReadFile("5.txt")

     if err != nil {
          log.Fatal(err)
     }

    input := string(content)

    // one(input)
    two(input)
}