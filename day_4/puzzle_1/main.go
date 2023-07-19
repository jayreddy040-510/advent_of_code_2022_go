package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func lineHandler(a, b string) bool {
	sliceA := strings.Split(a, "-")
	sliceB := strings.Split(b, "-")

	lowerA, err := strconv.Atoi(sliceA[0])
	if err != nil {
		log.Fatal(err)
	}

	lowerB, err := strconv.Atoi(sliceB[0])
	if err != nil {
		log.Fatal(err)
	}

	upperA, err := strconv.Atoi(sliceA[1])
	if err != nil {
		log.Fatal(err)
	}

	upperB, err := strconv.Atoi(sliceB[1])
	if err != nil {
		log.Fatal(err)
	}

	if lowerA <= lowerB && upperA >= upperB {
		return true
	} else if lowerB <= lowerA && upperB >= upperA {
		return true
	} else {
		return false
	}

}

func main() {
	f, err := os.Open("../puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(f)
	count := 0

	for {
		line, err := reader.ReadString('\n')
		line = strings.Trim(line, "\n")
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if lineHandler(strings.Split(line, ",")[0], strings.Split(line, ",")[1]) {
			count++
		}
	}
	fmt.Println(count)
}
