package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	title, _ := regexp.Compile("^\\d+:$")
	query, _ := regexp.Compile("^(\\d+x\\d+):\\s(.+)$")
	scanner := bufio.NewScanner(file)
	blocks := make([]int, 0)
	ans1 := 0
	for scanner.Scan() {
		line := scanner.Text()
		if title.MatchString(line) {
			blocks = append(blocks, 0)
		} else if query.MatchString(line) {
			parts := query.FindStringSubmatch(line)
			sizes := strings.Split(parts[1], "x")
			h, _ := strconv.Atoi(sizes[0])
			w, _ := strconv.Atoi(sizes[1])
			fills := strings.Split(parts[2], " ")
			needed_block := 0
			for i := 0; i < len(blocks); i++ {
				count, _ := strconv.Atoi(fills[i])
				needed_block += blocks[i] * count
			}
			if h*w >= needed_block {
				ans1 += 1
			}

		} else {
			// count number of # in line
			c := strings.Count(line, "#")
			blocks[len(blocks)-1] += c
		}

	}
	fmt.Println(blocks)
	fmt.Println(ans1)
}
