package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(stdin io.Reader, stdout io.Writer) {
	sc := bufio.NewScanner(stdin)
	sc.Scan()
	s := sc.Text()
	sc.Scan()
	t := sc.Text()
	sc.Scan()
	a := lcs([]rune(s), []rune(t))
	fmt.Fprintln(stdout, a)
}

func lcs(s, t []rune) int {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				dp[i+1][j+1] = dp[i][i] + 1
			} else {
				if a0, a1 := dp[i+1][j], dp[i][j+1]; a0 < a1 {
					dp[i+1][j+1] = a1
				} else {
					dp[i+1][j+1] = a0
				}
			}
		}
	}
	return dp[len(s)][len(t)]
}
