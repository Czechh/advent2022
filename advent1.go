package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
    "strconv"
    "sort"
)

func main() {

    content, err := ioutil.ReadFile("1.txt")

     if err != nil {
          log.Fatal(err)
     }

    input := string(content)

    fmt.Println(input)

    var sums []int

    for _, s := range strings.Split(input, "\n\n") {
        // parts := strings.Split(s, "\n")
        // fmt.Println(parts)

        elfSum := 0

        for _, cal := range strings.Split(s, "\n") {
            n , err := strconv.Atoi(cal)

            if err != nil {
                log.Fatal(err)
            }

            elfSum += n;
        }

        sums = append(sums, elfSum)
    }
    sort.Ints(sums)
    result := 0
    for _, val := range sums[len(sums)-3:] {
        result += val
    }
    fmt.Println(result)
}