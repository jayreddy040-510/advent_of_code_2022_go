package main

import (
	"fmt"
    "os"
	"strconv"
	"strings"
    "sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func maxGroupSum(filename string) int {
	data, err := os.ReadFile(filename)
	check(err)

	// Convert byte slice to string
	dataString := string(data)

	// Split the string into groups
	groups := strings.Split(dataString, "\n\n")

	// Initialize maximum sum
	maxSum := 0
    
    // Initialize slice of group counts
    var allSums []int

	// Iterate over the groups
	for _, group := range groups {
		// Split the group into numbers
		numbers := strings.Split(strings.Trim(group, "\n"), "\n")

		// Initialize group sum
		groupSum := 0

		// Iterate over the numbers and add them to the group sum
		for _, numStr := range numbers {
			num, err := strconv.Atoi(numStr)
			check(err)
			groupSum += num
		}

        allSums = append(allSums, groupSum)

		// Update maximum sum if necessary
		if groupSum > maxSum {
			maxSum = groupSum
		}
	}
    sort.Ints(allSums)

    return allSums[len(allSums) - 1] + allSums[len(allSums) - 2] + allSums[len(allSums) - 3]
}

func main() {
	filename := "elves_and_their_calories.txt"
	fmt.Printf("The sorted sums are %v\n", maxGroupSum(filename))
}


