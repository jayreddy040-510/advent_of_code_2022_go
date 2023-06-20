package main

import (
    "fmt"
    "strings"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func makePrioMaps() (map[rune]int) {

    prioMap := make(map[rune]int)
    
    x := 'a'
    y := 1

    for y <= 26 {
        prioMap[x] = y
        x += 1
        y += 1
    }

    x = 'A'

    for y <= 52 {
        prioMap[x] = y
        x += 1
        y += 1
    }

    return prioMap
}

func main() {

    prioMap := makePrioMaps()
    returnSum := 0
    data, err := os.ReadFile("puzzle_input.txt")
    check(err)

    dataString := string(data)

    rucksackSlice := strings.Split(dataString, "\n")

    for _, sack := range rucksackSlice {
        firstCompartmentContents := make(map[rune]bool)
        for i, char := range sack {
            if i < len(sack)/2 {
                firstCompartmentContents[char] = true
            } else {
                if firstCompartmentContents[char] {
                    // fmt.Println(string(char), prioMap[char])
                    returnSum += prioMap[char]
                    break
                }
            }
        }
    }

    fmt.Println(returnSum)

}
