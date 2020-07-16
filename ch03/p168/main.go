package main

import (
	"io"
	"log"
	"math"
	"os"
	"sort"

	"github.com/yuyamada/atcoder/lib"
)

var logger = log.New(os.Stderr, "", 0)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(stdin io.Reader, stdout io.Writer) {
	io := lib.NewIo(stdin, stdout)
	defer io.Flush()
	n, m := io.NextInt(), io.NextInt()
	a := io.NextInts(n)
	q := make([][]int, m)
	for i := 0; i < m; i++ {
		q[i] = io.NextInts(3)
	}
	ans := solver(a, q)
	for i := range ans {
		io.Println(ans[i])
	}
}

func solver(a []int, q [][]int) []int {
	n, m := len(a), len(q)
	ans := make([]int, m)
	sd := newSqrtDecomposition(a)

	sort.Ints(a)
	for i := range q {
		begin, end, k := q[i][0]-1, q[i][1], q[i][2]
		j := sort.Search(n, func(j int) bool { return sd.get(begin, end, a[j]) >= k })
		ans[i] = a[j]
	}
	return ans
}

type sqrtDecomposition struct {
	Data   []int
	Bucket [][]int
}

func newSqrtDecomposition(data []int) *sqrtDecomposition {
	n := len(data)
	cp := make([]int, n)
	for i := range data {
		cp[i] = data[i]
	}
	sqrtN := int(math.Sqrt(float64(n)))
	bucket := make([][]int, n/sqrtN+1)
	for i := 0; i < n; i++ {
		bucket[i/sqrtN] = append(bucket[i/sqrtN], data[i])
	}
	for i := range bucket {
		sort.Ints(bucket[i])
	}
	return &sqrtDecomposition{
		Data:   cp,
		Bucket: bucket,
	}
}

func (s *sqrtDecomposition) get(begin, end, x int) (count int) {
	n := len(s.Data)
	sqrtN := int(math.Sqrt(float64(n)))
	for i := range s.Bucket {
		l, r := i*sqrtN, (i+1)*sqrtN
		if r <= begin || end <= l {
			continue
		}
		if begin <= l && r <= end {
			j := sort.Search(len(s.Bucket[i]), func(j int) bool { return s.Bucket[i][j] > x })
			count += j
		} else {
			for j := lib.Max(begin, l); j < lib.Min(end, r); j++ {
				if s.Data[j] <= x {
					count++
				}
			}
		}
	}
	return count
}
