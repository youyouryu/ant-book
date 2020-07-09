package main

import (
	"io"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

const (
	qAdd = iota + 1
	qSum
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
		if query == qAdd {
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
	ans := []int{}
	data := make([]interface{}, len(array))
	for i := range array {
		data[i] = array[i]
	}
	st := lib.NewSegmentTree(data, merge)
	for _, v := range queries {
		switch v[0] {
		case qAdd:
			begin, end, value := v[1], v[2], v[3]
			add(st, begin, end, 0, value, 0, st.Size-st.Offset)
		case qSum:
			begin, end := v[1], v[2]
			ans = append(ans, st.Calc(begin, end).(int))
		}
	}
	return ans
}

func merge(a, b interface{}) interface{} {
	return a.(int) + b.(int)
}

func add(st *lib.SegmentTree, begin, end, k, value, cBegin, cEnd int) {
	if cEnd <= begin || end <= cBegin {
		return
	} else if begin <= cBegin && cEnd <= end {
		st.Nodes[k] = st.Nodes[k].(int) + value*(cEnd-cBegin)
	} else {
		st.Nodes[k] = st.Nodes[k].(int) + value*(lib.Min(end, cEnd)-lib.Max(begin, cBegin))
		add(st, begin, end, 2*k+1, value, cBegin, (cBegin+cEnd)/2)
		add(st, begin, end, 2*k+2, value, (cBegin+cEnd)/2, cEnd)
	}
}
