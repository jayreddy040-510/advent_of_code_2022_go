package main

import (
<<<<<<< HEAD
	"fmt"
	"os"
=======
	"bufio"
	"fmt"
	"os"
	"strconv"
>>>>>>> 959adc6 (fixes first puzzle, needed to clean string before Atoi conversion and wasn't printf'ing Atoi error.)
	"strings"
)

func main() {
<<<<<<< HEAD
	f, err := os.ReadFile("../elves_and_their_calories.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("error reading file: %v", err))
	}

	fmt.Println(string(f))
	fmt.Println("--------------------------------------------")
	fmt.Println(strings.Split(string(f), "\n\n"))
	// s := bufio.NewScanner(f)

	// for s.Scan() {
	// 	fmt.Println(s.Text())
	// }
}
=======
	f, err := os.Open("../elves_and_their_calories.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	line, err := reader.ReadString('\n')
	sum, maxSum := 0, 0
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
	}
	for err == nil {
		// fmt.Printf("Line: %v\n", line)
		if line == "" || line == "\n" {
			if sum > maxSum {
				maxSum = sum
			}
			sum = 0
			line, err = reader.ReadString('\n')
			continue
		}
		line = strings.TrimSuffix(line, "\n")
		if num, ok := strconv.Atoi(line); ok == nil {
			fmt.Printf("Num: %v\n", num)
			sum += num
			fmt.Printf("Sum: %v\n", sum)
		} else {
			fmt.Printf("error converting string to int: %v\n", ok)
		}
		line, err = reader.ReadString('\n')
	}
	fmt.Printf("Max sum: %v", maxSum)

}
>>>>>>> 959adc6 (fixes first puzzle, needed to clean string before Atoi conversion and wasn't printf'ing Atoi error.)
