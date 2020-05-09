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
	a := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		a = append(a, ai)
	}
	// ans := search(n, m, a)
	ans := search2(n, m, a)
	fmt.Fprintln(stdout, ans)
}

func search(n, m int, a []int) int {
	/*
	  すべて等価な個数制限付きナップザック問題として解く
	  dp[i][j] = i種類目までを使ったj個の組み合わせ総数
	*/
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := 0; j < n; j++ {
			if j == 0 || (i == 0 && j <= a[i]) {
				dp[i][j] = 1
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= m; j++ {
			for k := 0; k <= j && k <= a[i]; k++ {
				dp[i][j] += dp[i-1][j-k]
			}
		}
	}
	return dp[n-1][m]
}

func search2(n, m int, a []int) int {
	/*
	   kのループを除去する
	*/
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := 0; j < n; j++ {
			if j == 0 || (i == 0 && j <= a[i]) {
				dp[i][j] = 1
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = dp[i-1][j] + dp[i-1][j]
			if j-a[i]-1 >= 0 {
				dp[i][j] -= dp[i-1][j-a[i]-1]
			}
		}
	}
	return dp[n-1][m]
}
