package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	min, max int
}

func main() {
	path := "input.txt"

	byteFileContent, _ := readFile(path)
	sliceOfLines := strings.Split(string(byteFileContent), "\n")
	sliceOfRanges := make([]string, 0)
	sliceOfMergedRanges := make([]Range, 0) // format min-max just like the input
	sliceOfMergedRanges = append(sliceOfMergedRanges, Range{min: 0, max: 0})

	for _, v := range sliceOfLines {
		v = strings.TrimSpace(v)
		if strings.Contains(v, "-") {
			sliceOfRanges = append(sliceOfRanges, v)
		}
	}

	for _, v := range sliceOfRanges {
		valueMin, _ := strconv.Atoi(strings.Split(v, "-")[0])
		valueMax, _ := strconv.Atoi(strings.Split(v, "-")[1])

		for _, r := range sliceOfMergedRanges {
			if valueMax > r.max && valueMin >= r.min && valueMin <= r.max+1 {
				r.max = valueMax
				break
			}

			if valueMin < r.min && valueMax <= r.max && valueMax >= r.min-1 {
				r.min = valueMin
				break
			}

			if valueMin < r.min && valueMax > r.max {
				r.min = valueMin
				r.max = valueMax
				break
			}

			if valueMax < r.min-1 {
				sliceOfMergedRanges = append([]Range{{min: valueMin, max: valueMax}}, sliceOfMergedRanges...)
				break
			}

			if valueMin > r.max+1 {
				sliceOfMergedRanges = append(sliceOfMergedRanges, Range{min: valueMin, max: valueMax})
				break
			}
		}
	}

	for _, v := range sliceOfMergedRanges {
		fmt.Printf("Range: %v \n", v)
	}
}

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
