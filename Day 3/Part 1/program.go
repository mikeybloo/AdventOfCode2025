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
		firstDigit, firstDigitIndex := 0, 0
		fmt.Printf("Line: %v \n", sliceOfLines[i])

		for n := 0; n < len(sliceOfLines[i])-2; n++ {
			convertedDigit, _ := strconv.Atoi(string(sliceOfLines[i][n]))
			if convertedDigit > firstDigit {
				firstDigit = convertedDigit
				firstDigitIndex = n
			}
		}

		secondDigit := 0
		for j := len(sliceOfLines[i]) - 1; j > firstDigitIndex; j-- {
			convertedDigit, _ := strconv.Atoi(string(sliceOfLines[i][j]))
			if convertedDigit > secondDigit {
				secondDigit = convertedDigit
			}
		}

		fmt.Printf("Aquired batteries: %v \n", (firstDigit*10)+secondDigit)

		sum += (firstDigit * 10) + secondDigit
	}

	fmt.Printf("The sum of batteries: %v", sum)
}

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
