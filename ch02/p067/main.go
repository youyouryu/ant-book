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
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= m; j++ {
			for k := 0; k <= j && k <= a[i]; k++ {
				dp[i+1][j] += dp[i][j-k]
			}
		}
	}
	for i := range dp {
		fmt.Println(dp[i])
	}
	return dp[n][m]
}

func search2(n, m int, a []int) int {
	/*
	   kのループを除去する
	*/
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= m; j++ {
			dp[i+1][j] = dp[i+1][j-1] + dp[i][j]
			if j-1-a[i] >= 0 {
				dp[i+1][j] -= dp[i][j-1-a[i]]
			}
		}
	}
	return dp[n][m]
}
