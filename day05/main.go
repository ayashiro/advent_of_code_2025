package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

type Ranges []Range

func (r Ranges) Len() int {
	return len(r)
}

func (r Ranges) Less(i, j int) bool {
	if r[i].start != r[j].start {
		return r[i].start < r[j].start
	}
	return r[i].end < r[j].end
}

func (r Ranges) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func checkInclude(r Range, n int) bool {
	return r.start <= n && n <= r.end
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	ranges := make(Ranges, 0)
	numbers := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, Range{start, end})
	}
	ans1 := 0
	for scanner.Scan() {
		line := scanner.Text()
		i, _ := strconv.Atoi(line)
		numbers = append(numbers, i)
		for _, r := range ranges {
			if checkInclude(r, i) {
				ans1 += 1
				break
			}
		}
	}
	sort.Sort(ranges)
	mergedRange := make(Ranges, 0)
	for _, r := range ranges {
		if len(mergedRange) == 0 {
			mergedRange = append(mergedRange, r)
		} else {
			last := mergedRange[len(mergedRange)-1]
			if r.start <= last.end {
				mergedRange[len(mergedRange)-1] = Range{last.start, max(last.end, r.end)}
			} else {
				mergedRange = append(mergedRange, r)
			}
		}
	}

	fmt.Println(ans1)
	ans2 := 0
	for _, n := range mergedRange {
		ans2 += n.end - n.start + 1
	}
	fmt.Println(ans2)
}
