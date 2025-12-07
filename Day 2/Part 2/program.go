package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//insanely unoptimized but it works!
	path := "input.txt"

	byteFileContent, _ := readFile(path)
	sliceOfRanges := strings.Split(string(byteFileContent), ",")
	sliceOfNumbers := make([]string, 0)
	mapOfInvalidIds := make(map[string]bool)

	sum := 0

	for i := 0; i < len(sliceOfRanges); i++ {
		minN, _ := strconv.Atoi(strings.Split(sliceOfRanges[i], "-")[0])
		maxN, _ := strconv.Atoi(strings.Split(sliceOfRanges[i], "-")[1])

		for j := minN; j <= maxN; j++ {
			convertedJ := strconv.Itoa(j)

			sliceOfNumbers = append(sliceOfNumbers, convertedJ)
		}
	}

	for _, num := range sliceOfNumbers {
		fmt.Println("Number: ", num)
		for size := 0; size < len(num)/2; size++ {
			if len(num)%(size+1) == 0 {
				comparedString := num[:size+1]
				isInvalid := true
				for i := 0; i+size < len(num); i = i + size + 1 {
					fmt.Printf("\t Checking slice %v \n", num[i:i+size+1])
					if num[i:i+size+1] != comparedString {
						isInvalid = false
						break
					}
				}

				if isInvalid {
					mapOfInvalidIds[num] = true
					fmt.Printf("FOUND INVALID ID: %v \n", num)
					break
				}
			}
		}
	}

	for k, _ := range mapOfInvalidIds {
		convertedStr, _ := strconv.Atoi(k)
		sum += convertedStr
	}

	fmt.Println(sum)
}

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
