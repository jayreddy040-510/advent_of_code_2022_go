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

	if lowerA >= lowerB && lowerA <= upperB {
		return true
	} else if lowerB >= lowerA && lowerB <= upperA {
		return true
	} else if upperA >= lowerB && upperA <= upperB {
		return true
	} else if upperB >= lowerA && upperB <= upperA {
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
	blah := 0

	for {
		line, err := reader.ReadString('\n')
		line = strings.Trim(line, "\n")
		x := strings.Split(line, ",")[0]
		y := strings.Split(line, ",")[1]

		if err != nil {
			if err == io.EOF {
				if lineHandler(x, y) {
					count++
				}
				break
			}
			log.Fatal(err)
		}
		blah++
		if lineHandler(x, y) {
			count++
		}
	}

	fmt.Println(count)
}
