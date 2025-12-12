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
	twoDimSlice := make([][]string, 0)
	sum := 0

	for _, v := range sliceOfLines {
		newSlice := strings.Fields(strings.TrimSpace(v))
		twoDimSlice = append(twoDimSlice, newSlice)
	}

	fmt.Printf("slice: %v \n", twoDimSlice[len(twoDimSlice)-1])
	for i, v := range twoDimSlice[len(twoDimSlice)-1] {
		if v != "" {
			result := 0
			if v == "*" {
				result += 1
			}

			for n := 0; n < len(twoDimSlice)-1; n++ {
				number, _ := strconv.Atoi(twoDimSlice[n][i])
				fmt.Printf("Performing %v on number %v \n", v, number)
				if v == "*" {
					result *= number
				}

				if v == "+" {
					result += number
				}
			}

			sum += result
		}
	}

	fmt.Printf("The sum: %v", sum)
}

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
