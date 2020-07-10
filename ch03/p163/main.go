package main

import (
	"io"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(stdin io.Reader, stdout io.Writer) {
	io := lib.NewIo(stdin, stdout)
	defer io.Flush()
	n := io.NextInt()
	array := io.NextInts(n)
	q := io.NextInt()
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		query, begin, end := io.NextInt(), io.NextInt(), io.NextInt()
		queries[i] = []int{query, begin, end}
		if query == 1 {
			value := io.NextInt()
			queries[i] = append(queries[i], value)
		}
	}
	ans := solver(array, queries)
	for i := range ans {
		io.Println(ans[i])
	}
}

func solver(array []int, queries [][]int) []int {
	initialize(len(array))
	for i := range array {
		add(i, i+1, array[i], 0, 0, len(array))
	}
	ans := []int{}
	for i := range queries {
		q := queries[i]
		if q[0] == 1 {
			add(q[1], q[2], q[3], 0, 0, len(array))
		} else {
			ans = append(ans, sum(q[1], q[2], 0, 0, len(array)))
		}
	}
	return ans
}

var stSum, stAdd []int

func initialize(n int) {
	depth := 0
	for 1<<depth < n {
		depth++
	}
	stSum = make([]int, 1<<(depth+1)-1)
	stAdd = make([]int, 1<<(depth+1)-1)
}

func add(a, b, x, k, l, r int) {
	if a <= l && r <= b {
		stAdd[k] += x
	} else if l < b && a < r {
		stSum[k] += (lib.Min(b, r) - lib.Max(a, l)) * x
		add(a, b, x, 2*k+1, l, (l+r)/2)
		add(a, b, x, 2*k+2, (l+r)/2, r)
	}
}

func sum(a, b, k, l, r int) int {
	if b <= l || r <= a {
		return 0
	} else if a <= l && r <= b {
		return stAdd[k]*(r-l) + stSum[k]
	} else {
		res := (lib.Min(b, r) - lib.Max(a, l)) * stAdd[k]
		res += sum(a, b, 2*k+1, l, (l+r)/2)
		res += sum(a, b, 2*k+2, (l+r)/2, r)
		return res
	}
}
