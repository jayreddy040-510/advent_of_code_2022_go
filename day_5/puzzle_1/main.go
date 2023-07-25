package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func moveHandler(move string, cratesPointer *map[string][]string) {
	moveSentenceArray := strings.Split(move, " ")
	moveNum := moveSentenceArray[1]
	from := moveSentenceArray[3]
	to := moveSentenceArray[5]
}