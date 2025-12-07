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
	sum := 0

	for i := 0; i < len(sliceOfLines); i++ {
		line := strings.TrimSpace(sliceOfLines[i])
		sliceOfDigits, limitIndex := make([]int, 12), -1
		fmt.Printf("Line: %v \n", line)

		for n := 12; n > 0; n-- {
			newIndex := 0
			for x := len(line) - n; x > limitIndex; x-- {

				digit, _ := strconv.Atoi(string(line[x]))

				if digit >= sliceOfDigits[12-n] {
					sliceOfDigits[12-n] = digit
					newIndex = x
				}
			}

			limitIndex = newIndex
		}

		twelveDigits := ""
		for _, value := range sliceOfDigits {
			twelveDigits += strconv.Itoa(value)
		}

		numberToBeAdded, _ := strconv.Atoi(twelveDigits)
		fmt.Println("Number to be added:", numberToBeAdded)
		sum += numberToBeAdded
	}

	fmt.Printf("The sum of batteries: %v", sum)
}

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
