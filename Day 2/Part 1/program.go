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
	sliceOfRanges := strings.Split(string(byteFileContent), ",")

	sum := 0

	for i := 0; i < len(sliceOfRanges); i++ {
		fmt.Println(sliceOfRanges[i])
		minN, _ := strconv.Atoi(strings.Split(sliceOfRanges[i], "-")[0])
		maxN, _ := strconv.Atoi(strings.Split(sliceOfRanges[i], "-")[1])

		for j := minN; j <= maxN; j++ {
			convertedJ := strconv.Itoa(j)

			if len(convertedJ)%2 == 0 {
				if convertedJ[:len(convertedJ)/2] == convertedJ[len(convertedJ)/2:] {
					sum += j
				}
			}
		}
	}

	fmt.Println(sum)
}

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
