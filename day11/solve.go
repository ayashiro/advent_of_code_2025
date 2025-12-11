package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	graph := make(map[string][]string)
	prev := make(map[string]int)
	for scanner.Scan() {
		part := strings.Split(scanner.Text(), ": ")
		v := strings.Split(part[1], " ")
		graph[part[0]] = v
		for _, v := range v {
			prev[v] += 1
		}
	}
	sorted := make([]string, 0)
	qu := make([]string, 0)
	for k, _ := range graph {
		if prev[k] == 0 {
			qu = append(qu, k)
		}
	}
	for len(qu) > 0 {
		now := qu[0]
		qu = qu[1:]
		sorted = append(sorted, now)
		for _, v := range graph[now] {
			prev[v] -= 1
			if prev[v] == 0 {
				qu = append(qu, v)
			}
		}
	}
	points := []string{"svr", "fft", "dac", "you"}
	dps := make(map[string]map[string]int)
	for _, point := range points {
		dps[point] = make(map[string]int)
		dps[point][point] = 1
	}

	for _, v := range sorted {
		for _, nxt := range graph[v] {
			for _, point := range points {
				dps[point][nxt] += dps[point][v]
			}
		}
	}
	fmt.Println(dps["you"]["out"])
	fmt.Println(dps["svr"]["fft"]*dps["fft"]["dac"]*dps["dac"]["out"] + dps["svr"]["dac"]*dps["dac"]["fft"]*dps["fft"]["out"])
}
