package main

import (
	"bufio"
	"log"
	"os"
)

func rpsHandler(rps string) int {
	var pts int
	me, opp := rps[2], rps[0]

	if me == 'X' {
		pts += 1
		if opp == 'A' {
			pts += 3
		} else if opp == 'C' {
			pts += 6
		}
	} else if me == 'Y' {
		pts += 2
		if opp == 'A' {
			pts += 6
		} else if opp == 'B' {
			pts += 3
		}
	} else if me == 'Z' {
		pts += 3
		if opp == 'B' {
			pts += 6
		} else if opp == 'C' {
			pts += 3
		}
	}
	return pts
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
