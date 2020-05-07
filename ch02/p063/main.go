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
	a := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		a = append(a, ai)
	}
	// ans := search(a)
	ans := search2(a)
	fmt.Fprintln(stdout, ans)
}

func search(a []int) (ans int) {
	// 最大増加部分列長 = 系列長 - 重複数 - 反転数
	exists := map[int]bool{}
	ans = len(a)
	for i := 0; i < len(a)-1; i++ {
		if _, ok := exists[a[i]]; ok {
			ans--
		} else {
			exists[a[i]] = true
		}
		if a[i] > a[i+1] {
			ans--
		}
	}
	return
}

func search2(a []int) (ans int) {
	// dp[i] = 末尾がa[i]のときの最長部分増加列
	dp := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		if i == 0 {
			dp[i] = 1
			continue
		}
		for j := 0; j < i; j++ {
			if a[j] >= a[i] {
				continue
			}
			if dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}
	ans = dp[len(a)-1]
	return
}
