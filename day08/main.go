package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Location struct {
	x int
	y int
	z int
}

type container struct {
	dist int
	i    int
	j    int
}

type UnionFind interface {
	root(int) int
	union(int, int)
	size(int) int
}

type UnionFindImpl struct {
	data []int
}

func (uf *UnionFindImpl) root(x int) int {
	if uf.data[x] < 0 {
		return x
	}
	uf.data[x] = uf.root(uf.data[x])
	return uf.data[x]
}

func (uf *UnionFindImpl) union(x, y int) {
	rootX := uf.root(x)
	rootY := uf.root(y)
	if rootX != rootY {
		uf.data[rootX] += uf.data[rootY]
		uf.data[rootY] = rootX
	}
}
func (uf *UnionFindImpl) size(x int) int {
	return -uf.data[uf.root(x)]
}

func UnionFindInit(n int) UnionFind {
	uf := UnionFindImpl{}
	uf.data = make([]int, n)
	for i := 0; i < n; i++ {
		uf.data[i] = -1
	}
	return &uf
}

func distance(l1 Location, l2 Location) int {
	sq := func(x int) int { return x * x }
	return sq(l1.x-l2.x) + sq(l1.y-l2.y) + sq(l1.z-l2.z)
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	locations := make([]Location, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		locations = append(locations, Location{x, y, z})
	}
	containers := make([]container, 0)
	for i, p1 := range locations {
		for j, p2 := range locations {
			if i == j {
				break
			}
			c := container{distance(p1, p2), i, j}
			containers = append(containers, c)
		}
	}
	sort.Slice(containers, func(i, j int) bool {
		return containers[i].dist < containers[j].dist
	})
	uf := UnionFindInit(len(locations))
	N, _ := strconv.Atoi(os.Args[2])
	for i := 0; i < len(containers); i++ {
		c := containers[i]
		uf.union(c.i, c.j)
		if i == N {
			visitedRoot := make(map[int]bool)
			sizes := make([]int, 0)
			for i := 0; i < len(locations); i++ {
				root := uf.root(i)
				if visitedRoot[root] {
					continue
				}
				visitedRoot[root] = true
				sizes = append(sizes, uf.size(i))
			}
			sort.Ints(sizes)
			X := len(sizes)
			ans := sizes[X-1] * sizes[X-2] * sizes[X-3]
			fmt.Println(ans)
		}
		if uf.size(0) == len(locations) {
			fmt.Println(locations[c.i].x * locations[c.j].x)
			break
		}
	}

}
