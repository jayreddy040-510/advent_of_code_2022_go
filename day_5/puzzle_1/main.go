package main

import (
	"bufio"
	// "fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type stack struct {
	crates []string
}

func (s *stack) push(r string) {
	s.crates = append(s.crates, r)
}



func (s *stack) pop() string {
	lastItem := s.crates[len(s.crates)-1]
	s.crates = s.crates[:len(s.crates)-1]
	return lastItem
}

func moveHandler(crates *map[string]stack, move string) map[string]stack {
	moveSentenceArray := strings.Split(move, " ")
	// log.Println(moveSentenceArray[1], moveSentenceArray[3], moveSentenceArray[5])
	moveNum, err := strconv.Atoi(moveSentenceArray[1])

	if err != nil {
		log.Fatal(err)
	}

	from := moveSentenceArray[3]
	to := moveSentenceArray[5]

	for i := 0; i < moveNum; i++ {
		z := crates[from].pop()
		crates[to].push(z)
	}
	return crates
}

func main() {

	cratesMap := map[string]stack{
		"1": {crates: []string{"F", "T", "C", "L", "R", "P", "G", "Q"}},
		"2": {crates: []string{"N", "Q", "H", "W", "R", "F", "S", "J"}},
		"3": {crates: []string{"F", "B", "H", "W", "P", "M", "Q"}},
		"4": {crates: []string{"V", "S", "T", "D", "F"}},
		"5": {crates: []string{"Q", "L", "D", "W", "V", "F", "Z"}},
		"6": {crates: []string{"Z", "C", "L", "S"}},
		"7": {crates: []string{"Z", "B", "M", "V", "D", "F"}},
		"8": {crates: []string{"T", "J", "B"}},
		"9": {crates: []string{"Q", "N", "B", "G", "L", "S", "P", "H"}},
	}

	f, err := os.Open("../puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	// log.Println(cratesMap)

	for sc.Scan() {
		cratesMap = moveHandler(sc.Text(), cratesMap)
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	returnSlice := []string{}

	for i := 1; i < 10; i++ {
		iAscii := strconv.Itoa(i)

		theStack := cratesMap[iAscii].crates
		returnSlice = append(returnSlice, theStack[len(theStack)-1])
	}

	log.Println(strings.Join(returnSlice, ""))
}
