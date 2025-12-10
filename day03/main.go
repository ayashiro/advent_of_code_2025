package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	L := 12
	ans1 := 0
	ans2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		N := len(line)
		dp := make([][]int, L)
		for i := 0; i < L; i++ {
			dp[i] = make([]int, N)
		}
		elem := make([]int, 0)
		for _, c := range line {
			elem = append(elem, int(c-'0'))
		}
		for i := 0; i < L; i++ {
			for j := 0; j < N; j++ {
				if i == 0 {
					dp[i][j] = elem[j]
					if j > 0 {
						dp[i][j] = max(elem[j], dp[i][j-1])
					}
				} else if j > 0 {
					dp[i][j] = max(dp[i-1][j-1]*10+elem[j], dp[i][j-1])
				}
			}
		}
		ans1 += dp[1][N-1]
		ans2 += dp[11][N-1]
	}
	fmt.Println(ans1)
	fmt.Println(ans2)

}
