package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
    "strconv"
)

func one(input string) {
    result := 0

    for _, line := range strings.Split(input, "\n") {
        pairs := strings.Split(line, ",")

        pair0 := strings.Split(pairs[0], "-")
        pair1 := strings.Split(pairs[1], "-")

        p00, _ := strconv.Atoi(pair0[0])
        p01, _ := strconv.Atoi(pair0[1])

        p10, _ := strconv.Atoi(pair1[0])
        p11, _ := strconv.Atoi(pair1[1])

        fmt.Println(p00, p01, p10, p11)

        if p00 <= p10 && p01 >= p11 {
            result += 1
        } else if p10 <= p00 && p11 >= p01 {
            result += 1
        }
    }

    fmt.Println(result)
}

func two(input string) {
    result := 0

    for _, line := range strings.Split(input, "\n") {
        pairs := strings.Split(line, ",")

        pair0 := strings.Split(pairs[0], "-")
        pair1 := strings.Split(pairs[1], "-")

        p00, _ := strconv.Atoi(pair0[0])
        p01, _ := strconv.Atoi(pair0[1])

        p10, _ := strconv.Atoi(pair1[0])
        p11, _ := strconv.Atoi(pair1[1])

        fmt.Println(p00, p01, p10, p11)

        if p00 <= p10 && p01 >= p11 {
            result += 1
        } else if p10 <= p00 && p11 >= p01 {
            result += 1
        } else if p01 >= p10 && p00 <= p11 {
            result += 1
        } else if p11 >= p00 && p10 <= p01 {
            result += 1
        }
    }

    fmt.Println(result)
}

func main() {
    content, err := ioutil.ReadFile("4.txt")

     if err != nil {
          log.Fatal(err)
     }

    input := string(content)

    // one(input)
    two(input)
}