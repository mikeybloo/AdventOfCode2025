package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Feels a bit slow. Maybe could be improved...
func main() {
	path := "input.txt"

	byteFileContent, _ := readFile(path)
	sliceOfLines := strings.Split(string(byteFileContent), "\n")
	sliceOfRanges := make([]string, 0)
	sliceOfValues := make([]string, 0)

	count := 0

	for _, v := range sliceOfLines {
		v = strings.TrimSpace(v)
		if strings.Contains(v, "-") {
			sliceOfRanges = append(sliceOfRanges, v)
		} else if v != "" {
			sliceOfValues = append(sliceOfValues, v)
		}
	}

	for _, v := range sliceOfValues {
		fresh := false
		intV, _ := strconv.Atoi(v)
		fmt.Printf("Checking value %v ...\n", v)
		for _, r := range sliceOfRanges {
			min, _ := strconv.Atoi(strings.Split(r, "-")[0])
			max, _ := strconv.Atoi(strings.Split(r, "-")[1])
			fmt.Printf("\t Checking between range of %v and %v \n", min, max)

			if intV >= min && intV <= max {
				fmt.Printf("\t %v is fresh! \n", v)
				fresh = true
				break
			}
		}

		if fresh {
			count += 1
		} else {
			fmt.Printf("\t %v is rotten! \n", v)
		}
	}

	fmt.Printf("The count: %v", count)
}

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
