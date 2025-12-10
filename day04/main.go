package main

import (
	"bufio"
	"fmt"
	"os"
)

type coord struct {
	x int
	y int
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	grid := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		N := len(line)
		elem := make([]int, N)
		for i := 0; i < N; i++ {
			c := line[i]
			if c == '.' {
				elem[i] = 0
			} else {
				elem[i] = 1
			}
		}
		grid = append(grid, elem)
	}
	H := len(grid)
	W := len(grid[0])
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	ans := 0
	ans2 := 0
	for loop := 0; ; loop++ {
		pos := 0
		coords := make([]coord, 0)
		for x := 0; x < H; x++ {
			for y := 0; y < W; y++ {
				count := 0
				for d := 0; d < 8; d++ {
					nx, ny := x+dx[d], y+dy[d]
					if nx < 0 || nx >= H || ny < 0 || ny >= W {
						continue
					}
					count += grid[nx][ny]
				}
				if count < 4 && grid[x][y] == 1 {
					pos++
					ans2 += 1
					if loop == 0 {
						ans += 1
					}
					coords = append(coords, coord{x, y})
				}
			}
		}
		for _, c := range coords {
			grid[c.x][c.y] = 0
		}
		if pos == 0 {
			break
		}
	}
	fmt.Println(ans)
	fmt.Println(ans2)
}
