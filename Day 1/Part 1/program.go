package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "input.txt"

	byteFileContent, _ := readFile(path)
	sliceOfLines := strings.Split(string(byteFileContent), "\n")
	dial := 50
	counter := 0
	for i := 0; i < len(sliceOfLines); i++ {
		line := sliceOfLines[i]
		direction := line[:1]
		numeral, _ := strconv.Atoi(line[1 : len(line)-1])
		if direction == "L" {
			dial -= numeral
		} else {
			dial += numeral
		}

		if dial < 0 {
			if dial%100 == 0 {
				dial = 0
			} else {
				dial = 100 + (dial % 100)
			}
		} else if dial > 99 {
			dial = dial % 100
		}

		if dial == 0 {
			counter += 1
		}
	}
	fmt.Printf("Counter is: %v \n", counter)
}

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
