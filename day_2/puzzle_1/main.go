package main

import (
	"bufio"
	"log"
	"os"
)

// dont modify willy-nilly
var lookupTable = map[string]int{
    "A X": 4,
    "C X": 7,
    "B X": 1,
    "A Y": 8,
    "B Y": 5,
    "C Y": 2,
    "A Z": 3,
    "B Z": 9,
    "C Z": 6,
}

func rpsHandler (rps string) int {
    pts, exists := lookupTable[rps]; if exists {
        return pts
    } else {
        return 0
    }
}

func main() {
    file, err := os.Open("../puzzle_input.txt")
    if err != nil {
        log.Fatalf("Error reading puzzle input file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var totalPts int

    for scanner.Scan() {
        line := scanner.Text()
        totalPts += rpsHandler(line)
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error scanning puzzle input: %v", err)
    }

    log.Println(totalPts)

}
