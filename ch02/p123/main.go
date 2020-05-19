package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(stdin io.Reader, stdout io.Writer) {
	sc := bufio.NewScanner(stdin)
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	p, _ := strconv.ParseFloat(sc.Text(), 64)
	sc.Scan()
	x, _ := strconv.Atoi(sc.Text())
	ans := search(m, x, p)
	fmt.Fprintln(stdout, ans)
}

func search(m, x int, p float64) (ans float64) {
	const target = 1000000
	dp := make([][]float64, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]float64, 1<<(m-i)+1)
	}
	state := x * (1 << m) / target
	dp[0][state] = 1.0
	for i := 0; i < m; i++ {
		for j := range dp[i] {
			if j == 0 {
				dp[i+1][0] += dp[i][j] * 1.0
				continue
			}
			if j == len(dp[i])-1 {
				dp[i+1][len(dp[i+1])-1] += dp[i][j] * 1.0
				continue
			}
			dp[i+1][j/2+1] += dp[i][j] * p
			dp[i+1][(j-1)/2] += dp[i][j] * (1 - p)
		}
	}
	return dp[m][1]
}
