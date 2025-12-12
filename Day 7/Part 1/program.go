package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func main() {
	path := "input.txt"

	byteFileContent, _ := readFile(path)
	sliceOfLines := strings.Split(string(byteFileContent), "\n")
	twoDimSlice := make([][]string, 0)
	sum := 0

	for _, v := range sliceOfLines {
		newSlice := strings.Split(strings.TrimSpace(v), "")
		twoDimSlice = append(twoDimSlice, newSlice)
	}

	for i := 0; i < len(twoDimSlice)-1; i++ {
		for j := 0; j < len(twoDimSlice[i]); j++ {
			if twoDimSlice[i][j] == "S" {
				twoDimSlice[i+1][j] = "|"
				continue
			}

			if twoDimSlice[i][j] == "|" {
				if twoDimSlice[i+1][j] == "^" {
					twoDimSlice[i+1][j-1], twoDimSlice[i+1][j+1] = "|", "|"
					sum += 1
				} else {
					twoDimSlice[i+1][j] = "|"
				}
				continue
			}
		}

		//Visualize pyramid
		ClearScreen()
		drawPyramid(twoDimSlice)
	}

	fmt.Println("Sum is: ", sum)
}

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func drawPyramid(pyramid [][]string) {
	for _, valY := range pyramid {
		for _, valX := range valY {
			fmt.Print(valX)
		}
		fmt.Print("\n")
	}

	time.Sleep(50 * time.Millisecond)
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
