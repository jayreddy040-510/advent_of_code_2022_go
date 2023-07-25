package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"strings"
	"strconv"
)

func moveHandler(move string, cratesPointer map[string][]string) {
	moveSentenceArray := strings.Split(move, " ")
	moveNum, err := strconv.Atoi(moveSentenceArray[1])

	if err != nil {
		log.Fatal(err)
	}

	from := moveSentenceArray[3]
	to := moveSentenceArray[5]

	temp := cratesPointer[from][len(cratesPointer[from]) - moveNum : len(cratesPointer[from])]
	cratesPointer[from] = cratesPointer[from][:len(cratesPointer[from]) - moveNum]
	cratesPointer[to] = append(cratesPointer[to], temp...)
}

func main() {
	
}