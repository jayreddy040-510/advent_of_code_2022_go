package main

import (
    "fmt"
    "os"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    
    scoreMap := make(map[string]int)
    
    scoreMap["A X"] = 4
    scoreMap["B Y"] = 5
    scoreMap["C Z"] = 6
    scoreMap["A Y"] = 8
    scoreMap["A Z"] = 3
    scoreMap["B X"] = 1
    scoreMap["B Z"] = 9
    scoreMap["C X"] = 7
    scoreMap["C Y"] = 2

    data, err := os.ReadFile("../puzzle_input.txt")

    check(err)

    dataString := string(data)

    sliceOfStrings := strings.Split(strings.Trim(dataString, "\n"), "\n")

    totalScore := 0

    for _, bout := range sliceOfStrings {
        totalScore += scoreMap[bout]
    }
    fmt.Printf("\n The projected score is: %d", totalScore)
}
