package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func repeat(n int, times int) (int, error) {
	arr := []string{}
	s := strconv.Itoa(n)
	for i := 0; i < times; i++ {
		arr = append(arr, s)
	}
	return strconv.Atoi(strings.Join(arr, ""))
}

func findLowerBound(n int) int {
	if n <= 10 {
		return 0
	}
	left, right := 0, n
	for right-left > 1 {
		mid := (left + right) / 2
		d, _ := repeat(mid, 2)
		if d > n {
			right = mid
		} else {
			left = mid
		}
	}
	return left
}
func main() {
	primes := []int{2, 3, 5, 7}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Split(line, ",")
	ans := 0
	ans2 := 0
	for _, part := range parts {
		elems := strings.Split(part, "-")
		start, _ := strconv.Atoi(elems[0])
		end, _ := strconv.Atoi(elems[1])
		set := map[int]bool{}
		for _, rep := range primes {

			left, right := 0, start

			for right-left > 1 {
				mid := (left + right) / 2
				d, err := repeat(mid, rep)
				if err != nil || d >= start {
					right = mid
				} else {
					left = mid
				}
			}
			lowestRepeated := right
			left, right = 0, end
			for right-left > 1 {
				mid := (left + right) / 2
				d, err := repeat(mid, rep)
				if err != nil || d > end {
					right = mid
				} else {
					left = mid
				}
			}
			highestRepeated := left
			for i := lowestRepeated; i <= highestRepeated; i++ {
				d, _ := repeat(i, rep)
				set[d] = true
				if rep == 2 {
					ans += d
				}
			}
		}
		for d := range set {
			ans2 += d
		}

	}
	fmt.Println(ans)
	fmt.Println(ans2)
}
