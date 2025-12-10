package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid := make([][]int, 0)
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	start := 0
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if i == 0 {
			for c := 0; c < len(line); c++ {
				if line[c] == 'S' {
					start = c
					break
				}
			}
		}
		grid = append(grid, make([]int, len(line)))
		for c := 0; c < len(line); c++ {
			if line[c] == '.' {
				grid[i][c] = 0
			} else if line[c] == '^' {
				grid[i][c] = 1
			}
		}
	}
	grid[0][start] = 2
	ans1 := 0
	prev := make([]int, len(grid[0]))
	prev[start] = 1
	for i := 1; i < len(grid); i++ {
		now := make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			if prev[j] != 0 {
				if grid[i][j] == 0 {
					now[j] += prev[j]
				} else {
					ans1 += 1
					if grid[i][j+1] != 1 {
						now[j+1] += prev[j]
					}
					if grid[i][j-1] != 1 {
						now[j-1] += prev[j]
					}
				}
			}
		}
		prev = now
	}
	fmt.Println(ans1)
	ans2 := 0
	for _, v := range prev {
		ans2 += v
	}
	fmt.Println(ans2)
}
