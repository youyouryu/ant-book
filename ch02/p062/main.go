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
	k, _ := strconv.Atoi(sc.Text())
	a := []int{}
	m := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		line := strings.Split(sc.Text(), " ")
		ai, _ := strconv.Atoi(line[0])
		mi, _ := strconv.Atoi(line[1])
		a = append(a, ai)
		m = append(m, mi)
	}
	// ans := search(n, k, a, m, 0)
	// ans := search2(n, k, a, m, 0)
	ans := search3(n, k, a, m, 0)
	if ans {
		fmt.Fprintln(stdout, "Yes")
	} else {
		fmt.Fprintln(stdout, "No")
	}
}

func search(n, k int, a, m []int, idx int) bool {
	if k == 0 {
		return true
	}
	if idx == n {
		return false
	}
	var ans bool
	for i := 0; i <= m[idx]; i++ {
		if k >= a[idx]*i {
			ans = search(n, k-a[idx]*i, a, m, idx+1)
		}
		if ans {
			return true
		}
	}
	return false
}

func search2(n, k int, a, m []int, idx int) bool {
	dp := make([][]bool, len(a)+1)
	for i := range dp {
		dp[i] = make([]bool, k+1)
		dp[i][0] = true
	}
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < k+1; j++ {
			for l := 0; l <= m[i]; l++ {
				if j >= l*a[i] {
					dp[i][j] = dp[i][j] || dp[i+1][j-l*a[i]]
				}
			}
		}
	}
	return dp[0][k]
}

func search3(n, k int, a, m []int, idx int) bool {
	// 漸化式はiはi-1にのみ依存するので前の状態さえ残しておけばok
	dp := make([]int, k+1)
	for j := range dp {
		dp[j] = -1
	}
	dp[0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j <= k; j++ {
			if dp[j] >= 0 {
				// i-1までに合計がkとなる組み合わせが見つかっている
				// -> 一つも使わずok
				dp[j] = m[i]
				continue
			}
			if j < a[i] || dp[j-a[i]] <= 0 {
				// a[i]を使っても合計kを作ることができない
				// or a[i]が余っていない
				dp[j] = -1
				continue
			}
			// a[i]の使用可能数を一つ減らす
			dp[j] = dp[j-a[i]] - 1
		}
		// fmt.Println(i, a[i], m[i], dp)
	}
	return dp[k] >= 0
}
