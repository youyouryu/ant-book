package main

import (
	"io"
	"log"
	"math"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

var logger = log.New(os.Stderr, "", 0)

const inf = math.MaxInt64

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(r io.Reader, w io.Writer) {
	io := lib.NewIo(r, w)
	defer io.Flush()
	n, m := io.NextInt(), io.NextInt()
	costs := make([]map[int]int, n)
	for i := range costs {
		costs[i] = map[int]int{}
	}
	for i := 0; i < m; i++ {
		from, to, cost := io.NextInt(), io.NextInt(), io.NextInt()
		costs[from][to] = cost
	}
	ans := solver(costs)
	io.Println(ans)
}

func solver(costs []map[int]int) int {
	logger.Println(costs)
	n := len(costs)
	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0

	for s := 0; s < 1<<n; s++ {
		for v := 0; v < n; v++ {
			if s == 0 {
				if cost, ok := costs[0][v]; ok {
					dp[s|1<<v][v] = cost
					logger.Printf("%05b -> %05b %d\n", s, s|1<<v, dp[s|1<<v][v])
				}
				continue
			}
			if s>>v&1 == 1 {
				continue
			}
			min := inf
			for u := 0; u < n; u++ {
				if s>>u&1 == 0 {
					continue
				}
				if dp[s][u] == inf {
					continue
				}
				if _, ok := costs[u][v]; !ok {
					continue
				}
				dist := dp[s][u] + costs[u][v]
				if dist < min {
					min = dist
				}
			}
			if min < inf {
				dp[s|1<<v][v] = min
				logger.Printf("%05b -> %05b %d\n", s, s|1<<v, dp[s|1<<v][v])
			}
		}
	}
	return dp[1<<n-1][0]
}
