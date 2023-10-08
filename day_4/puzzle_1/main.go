package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isOverlap(elves string) (bool, error) {
	twoElvesGroup := strings.Split(elves, ",")
	elfOne := strings.Split(twoElvesGroup[0], "-")
	elfTwo := strings.Split(twoElvesGroup[1], "-")
	elfOneLower, err := strconv.Atoi(elfOne[0])
	if err != nil {
		return false, fmt.Errorf("Error converting elf sections: %s into numbers: %v", elves, err)
	}
	elfOneUpper, err := strconv.Atoi(elfOne[1])
	if err != nil {
		return false, fmt.Errorf("Error converting elf sections: %s into numbers: %v", elves, err)
	}
	elfTwoLower, err := strconv.Atoi(elfTwo[0])
	if err != nil {
		return false, fmt.Errorf("Error converting elf sections: %s into numbers: %v", elves, err)
	}
	elfTwoUpper, err := strconv.Atoi(elfTwo[1])
	if err != nil {
		return false, fmt.Errorf("Error converting elf sections: %s into numbers: %v", elves, err)
	}
	return (elfOneLower <= elfTwoLower && elfOneUpper >= elfTwoUpper) || (elfTwoLower <= elfOneLower && elfTwoUpper >= elfOneUpper), nil
}

func main() {
	file, err := os.Open("../puzzle_input.txt")
	if err != nil {
		log.Fatalf("Error opening puzzle input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var countOverlaps int

	for scanner.Scan() {
		line := scanner.Text()
		overlappingSection, err := isOverlap(line)
		if err != nil {
			log.Fatalf("%v", err)
		}
        if overlappingSection {
            countOverlaps++
        }
	}

	if scanner.Err() != nil {
		log.Fatalf("Error scanning puzzle input: %v", err)
	}

	log.Printf("Num overlaps: %d", countOverlaps)

}
