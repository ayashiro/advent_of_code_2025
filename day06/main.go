package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ADD = 0
	MUL = 1
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	N := len(lines)
	M := len(lines[0])
	operators := make([]int, 0)
	for _, c := range strings.Split(lines[N-1], " ") {
		if c == "+" {
			operators = append(operators, ADD)
		} else if c == "*" {
			operators = append(operators, MUL)
		} else {
			continue
		}
	}
	horizontals := make([][]int, 0)
	for i := 0; i < N-1; i++ {
		now := make([]int, 0)
		splitted := strings.Split(lines[i], " ")
		for _, element := range splitted {
			if strings.ReplaceAll(element, " ", "") == "" {
				continue
			}
			val, _ := strconv.Atoi(element)
			now = append(now, val)
		}
		horizontals = append(horizontals, now)
	}
	ans1 := 0
	for j, op := range operators {
		var base int
		if op == ADD {
			base = 0
		} else {
			base = 1
		}
		for i := 0; i < N-1; i++ {
			v := horizontals[i][j]
			if op == ADD {
				base += v
			} else {
				base *= v
			}
		}
		ans1 += base
	}
	fmt.Println(ans1)

	verticals := make([][]int, 0)
	now := make([]int, 0)
	for i := 0; i < M; i++ {
		tmp := ""
		for j := 0; j < N-1; j++ {
			tmp += string(lines[j][i])
		}
		rest := strings.ReplaceAll(tmp, " ", "")
		if rest == "" {
			verticals = append(verticals, now)
			now = make([]int, 0)
			continue
		}
		val, _ := strconv.Atoi(rest)
		now = append(now, val)

	}
	verticals = append(verticals, now)
	ans2 := 0
	for i, op := range operators {
		var base int
		if op == ADD {
			base = 0
		} else {
			base = 1
		}
		for _, v := range verticals[i] {
			if op == ADD {
				base += v
			} else {
				base *= v
			}
		}
		ans2 += base
	}
	fmt.Println(ans2)

}
