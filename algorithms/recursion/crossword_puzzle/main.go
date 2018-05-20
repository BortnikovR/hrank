//https://www.hackerrank.com/challenges/crossword-puzzle/problem
package main

import (
	"fmt"
	"strings"
)

const (
	plus  = "+"
	minus = "-"
)

var found = false

func main() {
	grid := make([][]string, 0, 10)
	for i := 0; i < 10; i++ {
		var str string
		fmt.Scan(&str)
		grid = append(grid, strings.Split(str, ""))
	}
	var str string
	fmt.Scan(&str)
	words := strings.Split(str, ";")
	putWord(words, grid)
}

func putWord(words []string, grid [][]string) {
	if found {
		return
	}
	if len(words) == 0 {
		found = true
		printGrid(grid)
	}
	for i, w := range words {
		if newGrid, ok := putHor(w, copyGrid(grid)); ok {
			temp := make([]string, 0, len(words) - 1)
			temp = append(temp, words[:i]...)
			temp = append(temp, words[i+1:]...)
			putWord(temp, newGrid)
		}
		if newGrid, ok := putVer(w, copyGrid(grid)); ok {
			temp := make([]string, 0, len(words) - 1)
			temp = append(temp, words[:i]...)
			temp = append(temp, words[i+1:]...)
			putWord(temp, newGrid)
		}
	}
}

func putHor(w string, grid [][]string) ([][]string, bool) {
	for j := 0; j < 10; j++ {
		for i := 0; i < 10; i++ {
			if isHorStart(i, j, grid) {
				if matchHor(w, i, j, grid) {
					for n := 0; n < len(w); n++ {
						grid[i][j+n] = w[n:n+1]
					}
					return grid, true
				}
			}
		}
	}
	return nil, false
}

func isHorStart(i, j int, grid [][]string) bool {
	if grid[i][j] == minus && (j == 0 || grid[i][j-1] == plus) {
		return true
	}
	return false
}

func matchHor(w string, i, j int, grid [][]string) bool {
	n := 0
	for ; j < 10 && grid[i][j] != plus; j++ {
		if grid[i][j] == minus || (n < len(w) && grid[i][j] == w[n:n+1]) {
			n++
		}
	}
	return n == len(w)
}

func putVer(w string, grid [][]string) ([][]string, bool) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if isVerStart(i, j, grid) {
				if matchVer(w, i, j, grid) {
					for n := 0; n < len(w); n++ {
						grid[i+n][j] = w[n:n+1]
					}
					return grid, true
				}
			}
		}
	}
	return nil, false
}

func isVerStart(i, j int, grid [][]string) bool {
	if grid[i][j] != plus && (i == 0 || grid[i-1][j] == plus) {
		return true
	}
	return false
}

func matchVer(w string, i, j int, grid [][]string) bool {
	n := 0
	for ; i < 10 && grid[i][j] != plus; i++ {
		if grid[i][j] == minus || (n < len(w) && grid[i][j] == w[n:n+1]) {
			n++
		}
	}
	return n == len(w)
}

func printGrid(grid [][]string) {
	for _, s := range grid {
		for _, c := range s {
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func copyGrid(grid [][]string) [][]string {
	gridCopy := make([][]string, len(grid))
	for i := range grid {
		gridCopy[i] = make([]string, len(grid[i]))
		copy(gridCopy[i], grid[i])
	}
	return gridCopy
}
