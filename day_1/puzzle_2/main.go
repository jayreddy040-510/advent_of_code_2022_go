package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func updateMaxSums(maxSums *[]int, sum int) {
	if sum > (*maxSums)[2] {
		(*maxSums)[2] = sum
		if (*maxSums)[2] > (*maxSums)[1] {
			(*maxSums)[2], (*maxSums)[1] = (*maxSums)[1], (*maxSums)[2]
			if (*maxSums)[1] > (*maxSums)[0] {
				(*maxSums)[1], (*maxSums)[0] = (*maxSums)[0], (*maxSums)[1]
			}
		}
	}
}

func main() {
	f, err := os.Open("../elves_and_their_calories.txt")
	if err != nil {
		fmt.Printf("error opening file: %s\n", err)
		return
	}
	defer f.Close()

	var sum int
	maxSums := []int{0, 0, 0}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			updateMaxSums(&maxSums, sum)
			sum = 0
			continue
		}

		num, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			fmt.Printf("error converting string to int: %v\n", err)
			continue
		}
		sum += num
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %s\n", err)
		return
	}

	updateMaxSums(&maxSums, sum)

	fmt.Println(maxSums[0] + maxSums[1] + maxSums[2])
}
