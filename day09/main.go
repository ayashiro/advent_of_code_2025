package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
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
	points := make([]Point, 0)
	for scanner.Scan() {
		line := scanner.Text()
		part := strings.Split(line, ",")
		a, _ := strconv.Atoi(part[0])
		b, _ := strconv.Atoi(part[1])
		points = append(points, Point{a, b})
	}
	xset := map[int]bool{}
	yset := map[int]bool{}
	for _, p := range points {
		xset[p.x] = true
		yset[p.y] = true
	}
	X := make([]int, 0)
	Y := make([]int, 0)
	for v := range xset {
		X = append(X, v)
	}
	for v := range yset {
		Y = append(Y, v)
	}
	maxX, minX := X[0], X[0]
	maxY, minY := Y[0], Y[0]
	for _, v := range X {
		if v > maxX {
			maxX = v
		}
		if v < minX {
			minX = v
		}
	}
	for _, v := range Y {
		if v > maxY {
			maxY = v
		}
		if v < minY {
			minY = v
		}
	}
	X = append(X, maxX+1)
	X = append(X, minX-1)

	Y = append(Y, maxY+1)
	Y = append(Y, minY-1)
	sort.Ints(X)
	sort.Ints(Y)
	xmap := make(map[int]int)
	ymap := make(map[int]int)
	for i, v := range X {
		xmap[v] = i
	}
	for i, v := range Y {
		ymap[v] = i
	}

	N := len(points)

	prev := Point{
		xmap[points[N-1].x],
		ymap[points[N-1].y],
	}
	grid := make([][]int, len(X))
	for i := 0; i < len(X); i++ {
		grid[i] = make([]int, len(Y))
	}
	for i := 0; i < N; i++ {
		x1, y1 := prev.x, prev.y
		x2, y2 := xmap[points[i].x], ymap[points[i].y]
		grid[x2][y2] = 1
		if x1 == x2 {
			for j := min(y1, y2) + 1; j < max(y1, y2); j++ {
				grid[x1][j] = 2
			}
		}
		if y1 == y2 {
			for j := min(x1, x2) + 1; j < max(x1, x2); j++ {
				grid[j][y1] = 2
			}
		}
		prev = Point{x2, y2}
	}
	queue := make([]Point, 0)
	queue = append(queue, Point{0, 0})
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if grid[p.x][p.y] == -1 {
			continue
		}
		grid[p.x][p.y] = -1
		dx := []int{0, 1, 0, -1}
		dy := []int{1, 0, -1, 0}
		for i := 0; i < 4; i++ {
			x := p.x + dx[i]
			y := p.y + dy[i]
			if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
				continue
			}
			if grid[x][y] != 0 {
				continue
			}
			queue = append(queue, Point{x, y})
		}
	}
	H := len(X)
	W := len(Y)
	dp := make([][]int, H)
	for i := 0; i < H; i++ {
		dp[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		count := 0
		for j := 0; j < W; j++ {
			val := 0
			if i != 0 {
				val = dp[i-1][j]
			}
			if grid[i][j] >= 0 {
				count += 1
			}
			dp[i][j] = val + count
		}
	}
	ans1 := 0
	ans2 := 0
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			p1 := points[i]
			p2 := points[j]
			area := (abs(p1.x-p2.x) + 1) * (abs(p1.y-p2.y) + 1)
			maxX, minX = xmap[max(p1.x, p2.x)], xmap[min(p1.x, p2.x)]
			maxY, minY = ymap[max(p1.y, p2.y)], ymap[min(p1.y, p2.y)]
			area1 := (maxX - minX + 1) * (maxY - minY + 1)
			area2 := dp[maxX][maxY] - dp[maxX][minY-1] - dp[minX-1][maxY] + dp[minX-1][minY-1]
			ans1 = max(area, ans1)
			if area1 == area2 {
				ans2 = max(area, ans2)
			}
		}
	}
	fmt.Println(ans1)
	fmt.Println(ans2)

}
