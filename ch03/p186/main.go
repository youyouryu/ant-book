package main

import (
	"io"
	"log"
	"math"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(r io.Reader, w io.Writer) {
	io := lib.NewIo(r, w)
	defer io.Flush()
	n, m := io.NextInt(), io.NextInt()
	sorters := make([]sorter, m)
	for i := 0; i < m; i++ {
		first, last := io.NextInt(), io.NextInt()
		sorters[i].first, sorters[i].last = first, last
	}
	ans := solver(n, sorters)
	io.Println(ans)
}

// O(mlogn)
func solver(n int, sorters []sorter) (ans int) {
	dp := make([]interface{}, n+1)
	for j := range dp {
		dp[j] = inf
	}
	dp[1] = 0
	st := lib.NewSegmentTree(dp, merge)
	m := len(sorters)
	for i := 0; i < m; i++ {
		first, last := sorters[i].first, sorters[i].last
		// O(logn)
		min := st.Query(first, last+1).(int)
		if min < inf {
			// O(logn)
			st.Update(st.Offset+last, min+1)
		}
	}
	return st.Nodes[st.Offset+n].(int)
}

func merge(a, b interface{}) interface{} {
	return lib.Min(a.(int), b.(int))
}

// O(nm)
func solver2(n int, sorters []sorter) (ans int) {
	m := len(sorters)
	logger.Println(n, sorters)
	dp := make([]int, n+1)
	for j := range dp {
		dp[j] = inf
	}
	dp[1] = 0
	for i := 0; i < m; i++ {
		first, last := sorters[i].first, sorters[i].last
		min := lib.Min(dp[first : last+1]...)
		if min < inf {
			dp[last] = min + 1
		}
	}
	return dp[n]
}

const inf = math.MaxInt64

var logger = log.New(os.Stderr, "", 0)

type sorter struct{ first, last int }
