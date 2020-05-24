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
	s, _ := strconv.Atoi(sc.Text())
	a := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		a = append(a, ai)
	}
	ans := search(s, a)
	fmt.Fprintln(stdout, ans)
}

func search(s int, a []int) (ans int) {
	n := len(a)
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j+i >= n {
				continue
			}
			dp[j] = dp[j+1] + a[j]
			if dp[j] >= s {
				return i + 1
			}
		}
	}
	return -1
}
