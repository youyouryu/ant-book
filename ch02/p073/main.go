package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(stdin io.Reader, stdout io.Writer) {
	sc := bufio.NewScanner(stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	l, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	p, _ := strconv.Atoi(sc.Text())
	a, b := []int{}, []int{}
	for sc.Scan() {
		line := strings.Split(sc.Text(), " ")
		ai, _ := strconv.Atoi(line[0])
		bi, _ := strconv.Atoi(line[1])
		a = append(a, ai)
		b = append(b, bi)
	}
	ans := search(n, l, p, a, b)
	fmt.Fprintln(stdout, ans)
}

func search(n, l, p int, a, b []int) int {
	/*
	  place := {S, A1, A2, ..., An, G}
	  dp[i][j] := place[i]までの補給回数がj回のときのガソリン残量
	  計算量：O(n^2)
	*/
	// dpテーブルの初期化
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[0][0] = p

	// 探索
	a = append([]int{0}, append(a, l)...)
	b = append([]int{0}, append(b, 0)...)
	for i := 0; i < n+2-1; i++ {
		for j := 0; j <= i && j < n-1; j++ {
			if rest := dp[i][j] - (a[i+1] - a[i]); rest < 0 {
				dp[i+1][j] = -1
			} else {
				dp[i+1][j] = rest
			}

			// 地点i+1で補給i{せず,して}合計j+1回
			var rest0, rest1 int
			if dp[i][j+1] < 0 {
				rest0 = -1
			} else {
				rest0 = dp[i][j+1] - (a[i+1] - a[i])
			}
			if dp[i+1][j] < 0 {
				rest1 = -1
			} else {
				rest1 = dp[i+1][j] + b[i+1]
			}
			if rest0 < rest1 {
				dp[i+1][j+1] = rest1
			} else {
				dp[i+1][j+1] = rest0
			}
		}
	}

	// 最小補給回数を返す
	for j := range dp[n+1] {
		if dp[n+1][j] >= 0 {
			return j
		}
	}
	return -1
}
