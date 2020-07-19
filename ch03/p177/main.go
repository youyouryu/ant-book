package main

import (
	"io"
	"log"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(r io.Reader, w io.Writer) {
	io := lib.NewIo(r, w)
	defer io.Flush()
	n, _ := io.NextInt(), io.NextInt()
	stage := make([]string, n)
	for i := 0; i < n; i++ {
		stage[i] = io.Next()
	}
	ans := solver(stage)
	io.Println(ans)
}

func solver(stage []string) (ans int) {
	for i := range stage {
		logger.Println(stage[i])
	}
	n, m := len(stage), len(stage[0])
	dp := make([][][][]int, 2)
	for i := range dp {
		// add a dummy row
		dp[i] = make([][][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([][]int, m)
			for k := range dp[i][j] {
				dp[i][j][k] = make([]int, 1<<m)
			}
		}
	}
	curr, next := dp[0], dp[1]
	curr[0][0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var iNext, jNext, sNext int
			if j+1 < m {
				iNext, jNext = i, j+1
			} else {
				iNext, jNext = i+1, 0
			}
			for s := 0; s < 1<<m; s++ {
				if s>>j&1 == 1 || stage[i][j] == 'x' {
					sNext = s &^ (1 << j)
					next[iNext][jNext][sNext] += curr[i][j][s]
					next[iNext][jNext][sNext] %= mod
					if curr[i][j][s] != 0 {
						logger.Printf("(%d %d) %05b %d =s=> (%d %d) %05b %d\n", i, j, s, curr[i][j][s], iNext, jNext, sNext, next[iNext][jNext][sNext])
					}
				} else {
					if j+1 < m && s>>(j+1)&1 == 0 && stage[i][j+1] == '.' {
						sNext = s | 1<<(j+1)
						next[iNext][jNext][sNext] += curr[i][j][s]
						next[iNext][jNext][sNext] %= mod
						if curr[i][j][s] != 0 {
							logger.Printf("(%d %d) %05b %d =h=> (%d %d) %05b %d\n", i, j, s, curr[i][j][s], iNext, jNext, sNext, next[iNext][jNext][sNext])
						}
					}
					if i+1 < n && stage[i+1][j] == '.' {
						sNext = s | 1<<j
						next[iNext][jNext][sNext] += curr[i][j][s]
						next[iNext][jNext][sNext] %= mod
						if curr[i][j][s] != 0 {
							logger.Printf("(%d %d) %05b %d =v=> (%d %d) %05b %d\n", i, j, s, curr[i][j][s], iNext, jNext, sNext, next[iNext][jNext][sNext])
						}
					}
				}
			}
			curr, next = next, curr
		}
	}
	return curr[n][0][0]
}

var mod int = 1e9 + 7
var logger = log.New(os.Stderr, "", 0)
