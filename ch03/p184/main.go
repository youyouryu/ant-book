package main

import (
	"io"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

var mod int

func solve(r io.Reader, w io.Writer) {
	io := lib.NewIo(r, w)
	defer io.Flush()
	n, k := io.NextInt(), io.NextInt()
	mod = io.NextInt()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = io.NextInts(n)
	}
	ans := solver(a, k)
	for i := 0; i < n; i++ {
		io.Printf("%d", ans[i][0])
		for j := 1; j < n; j++ {
			io.Printf(" %d", ans[i][j])
		}
		io.Println()
	}
}

func solver(a [][]int, k int) (ans [][]int) {
	b := lib.Matrix(a)
	zero := lib.NewMatrix(len(a), len(a))
	eye := lib.NewIdentityMatrix(len(a))
	c := stack([][]lib.Matrix{{b, zero}, {eye, eye}})
	d := vStack([]lib.Matrix{eye, zero})
	e := c.Pow(k).Mul(d)
	f := make(lib.Matrix, len(a))
	for i := range f {
		f[i] = e[len(a):][i][:len(a)]
	}
	return f.Mul(b)
}

func stack(ms [][]lib.Matrix) lib.Matrix {
	rows := make([]lib.Matrix, len(ms))
	for i := range ms {
		rows[i] = hStack(ms[i])
	}
	return vStack(rows)
}

func vStack(ms []lib.Matrix) lib.Matrix {
	ret := [][]int{}
	for k := range ms {
		ret = append(ret, ms[k]...)
	}
	return ret
}

func hStack(ms []lib.Matrix) lib.Matrix {
	ret := make([][]int, len(ms))
	for k := range ms {
		for i := range ms[k] {
			ret[i] = append(ret[i], ms[k][i]...)
		}
	}
	return ret
}
