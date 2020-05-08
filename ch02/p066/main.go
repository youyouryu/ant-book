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
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	M, _ := strconv.Atoi(sc.Text())
	// a := search(n, m, M)
	a := search2(n, m, M)
	fmt.Fprintln(stdout, a)
}

func search(n, m, M int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if i <= 1 || j == 1 {
				dp[i][j] = 1
				continue
			}
			for k := 0; k <= i/j; k++ {
				dp[i][j] += dp[i-k*j][j-1]
			}
		}
	}
	return dp[n][m] % M
}

func search2(n, m, M int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if i <= 1 || j == 1 {
				dp[i][j] = 1
				continue
			}
			if i < j {
				continue
			}
			dp[i][j] = dp[i-j][j] + dp[i][j-1]
		}
	}
	return dp[n][m] % M
}
