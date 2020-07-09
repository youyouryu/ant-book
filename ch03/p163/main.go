package main

import (
	"fmt"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

const (
	qAdd = iota + 1
	qSum
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	array := io.NextInts(n)
	q := io.NextInt()
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		query, begin, end := io.NextInt(), io.NextInt(), io.NextInt()+1
		queries[i] = []int{query, begin, end}
		if query == qAdd {
			value := io.NextInt()
			queries[i] = append(queries[i], value)
		}
	}
	ans := solve(array, queries)
	for i := range ans {
		io.Println(ans[i])
	}
}

func solve(array []int, queries [][]int) []int {
	fmt.Println(array, queries)
	data := make([]interface{}, len(array))
	for i := range array {
		data[i] = array[i]
	}
	st := lib.NewSegmentTree(data, merge)
	fmt.Println(st)
	for _, v := range queries {
		switch v[0] {
		case qAdd:
		case qSum:
		}
	}
	ans := make([]int, len(queries))
	return ans
}

type node struct {
	value, extra int
}

func merge(a, b interface{}) interface{} {
	na, nb := a.(node), b.(node)
	return node{na.value + nb.value, na.extra + nb.extra}
}

func queryAdd(st *lib.SegmentTree, begin, end, extra int) {
	add(st, begin, end, extra, 0, 0, len(st.nodes)-st.offset)
}

func add(st *lib.SegmentTree, begin, end, k, extra, cBegin, cEnd int) {
	if cEnd <= begin || end <= cBegin {
		return
	} else if begin <= cBegin && cEnd <= end {
		nd := st.Get(k).(node)
		nd.extra += extra
		st.Update(k, nd)
	} else {
		add(st, begin, end, 2*k+1, extra, cBegin, (cBegin+cEnd)/2)
		add(st, begin, end, 2*k+2, extra, (cBegin+cEnd)/2, cEnd)
	}
}
