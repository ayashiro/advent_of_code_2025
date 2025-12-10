package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	now := 50
	ans := 0
	ans2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		steps, _ := strconv.Atoi(line[1:])
		next := 0
		if direction == 'L' {
			next = now - steps
		} else {
			next = now + steps
		}
		if next%100 == 0 {
			ans += 1
		}
		var a, b int
		if direction == 'L' {
			a = now - 1
			b = next
		} else {
			a = next
			b = now + 1
		}
		tmp := 0
		if a == 0 {
			tmp = 1 - b/100
		} else if b <= 0 {
			tmp = a/100 - b/100 + 1
		} else {
			tmp = a/100 - b/100
		}
		ans2 += tmp
		now = (next%100 + 100) % 100

	}
	fmt.Println(ans)
	fmt.Println(ans2)

}
