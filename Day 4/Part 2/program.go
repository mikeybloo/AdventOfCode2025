package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	path := "input.txt"

	byteFileContent, _ := readFile(path)
	sliceOfLines := strings.Split(string(byteFileContent), "\n")
	sum := 0
	add := 1

	for add > 0 {
		add = 0
		for i := 0; i < len(sliceOfLines); i++ {
			line := strings.TrimSpace(sliceOfLines[i])

			for n := 0; n < len(line); n++ {
				leftLimit := 1
				rightLimit := 1
				if string(line[n]) == "@" {
					if n == 0 {
						leftLimit = 0
					}

					if n == len(line)-1 {
						rightLimit = 0
					}

					topRow := "..."
					medRow := strings.TrimSpace(sliceOfLines[i][n-leftLimit : n+rightLimit+1])
					botRow := "..."

					if i > 0 {
						topRow = strings.TrimSpace(sliceOfLines[i-1][n-leftLimit : n+rightLimit+1])
					}

					if i < len(sliceOfLines)-1 {
						botRow = strings.TrimSpace(sliceOfLines[i+1][n-leftLimit : n+rightLimit+1])
					}

					count := 0
					for x := 0; x < len(medRow); x++ {
						if string(medRow[x]) == "@" {
							count += 1
						}

						if string(topRow[x]) == "@" {
							count += 1
						}

						if string(botRow[x]) == "@" {
							count += 1
						}
					}

					if count-1 < 4 {
						add += 1
						b := []byte(sliceOfLines[i])
						b[n] = '.'
						sliceOfLines[i] = string(b)
					}
				}
			}
		}

		sum += add
	}

	fmt.Println("Sum:", sum)
}

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
