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

func boutHandler(bout string) int {
    sum := 0

    winnerMap := map[byte]int {
        'A': 2,
        'B': 3,
        'C': 1,
    }

    drawMap := map[byte]int {
        'A': 1,
        'B': 2,
        'C': 3,
    }

    loserMap := map[byte]int {
        'A': 3,
        'B': 1,
        'C': 2,
    }
    
    last := bout[2]
    first := bout[0]

    switch last {
    case 'X':
        sum += loserMap[first]
    case 'Y':
        sum += 3 + drawMap[first]
    case 'Z':
        sum += 6 + winnerMap[first]
    }
    return sum
}

func main() {
    
    data, err := os.ReadFile("../puzzle_input.txt")

    check(err)

    dataString := string(data)

    sliceOfStrings := strings.Split(strings.Trim(dataString, "\n"), "\n")

    totalScore := 0

    for _, bout := range sliceOfStrings {
        totalScore += boutHandler(bout)
    }
    fmt.Printf("\n The projected score is: %d", totalScore)
}
