package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
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