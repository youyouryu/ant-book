package main

import (
	"io"
	"log"
	"math"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

var logger = log.New(os.Stderr, "", 0)

const inf = math.MaxFloat64

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(r io.Reader, w io.Writer) {
	io := lib.NewIo(r, w)
	defer io.Flush()
	n, m := io.NextInt(), io.NextInt()
	costs := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		costs[i] = make(map[int]int)
	}
	for i := 0; i < m; i++ {
		from, to, cost := io.NextInt()-1, io.NextInt()-1, io.NextInt()
		costs[from][to] = cost
		costs[to][from] = cost
	}
	start, goal := io.NextInt()-1, io.NextInt()-1
	t := io.NextInt()
	tickets := io.NextInts(t)
	ans := solver(costs, tickets, start, goal)
	io.Printf("%.3f", ans)
}

func solver(costs []map[int]int, tickets []int, start, goal int) (ans float64) {
	logger.Println(costs, tickets)
	n, t := len(costs), len(tickets)
	dp := make([][]float64, 1<<t)
	for i := range costs {
		dp[i] = make([]float64, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	dp[1<<t-1][start] = 0

	for s := 1<<t - 1; s >= 0; s-- {
		for v := 0; v < n; v++ {
			if dp[s][v] == inf {
				logger.Printf("%05b v=%d -> continue\n", s, v)
				continue
			}
			for u := 0; u < n; u++ {
				if _, ok := costs[v][u]; !ok {
					continue
				}
				for i := 0; i < t; i++ {
					if s>>i&1 == 0 {
						continue
					}
					dist := dp[s][v] + float64(costs[v][u])/float64(tickets[i])
					if dist < dp[s&^(1<<i)][u] {
						dp[s^(1<<i)][u] = dist
						logger.Printf("%05b v=%d -> %05b u=%d dist=%.2f\n", s, v, s&^(1<<i), u, dist)
					}
				}
			}
		}
	}

	ans = inf
	for s := 0; s < 1<<t; s++ {
		if dp[s][goal] < ans {
			ans = dp[s][goal]
		}
	}
	return ans
}
