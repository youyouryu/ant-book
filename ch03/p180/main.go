package main

import (
	"fmt"
	"io"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(r io.Reader, w io.Writer) {
	io := lib.NewIo(r, w)
	defer io.Flush()
	n := io.NextInt()
	ans := fibonacci(n)
	io.Println(ans)
}

func fibonacci(n int) int {
	a := newMatrix(2, 2)
	a[0][0], a[0][1] = 1, 1
	a[1][0], a[1][1] = 1, 0
	b := newMatrix(2, 1)
	b[0][0] = 1
	b[1][0] = 0
	c := a.pow(n).mul(b)
	return c[1][0]
}

type matrix [][]int

func newMatrix(n, m int) matrix {
	ret := make(matrix, n)
	for i := range ret {
		ret[i] = make([]int, m)
	}
	return ret
}

func newIdentityMatrix(n int) matrix {
	ret := newMatrix(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				ret[i][j] = 1
			}
		}
	}
	return ret
}

func (a matrix) mul(b matrix) matrix {
	if len(a[0]) != len(b) {
		panic(fmt.Errorf("shape mismatch"))
	}
	n, m, l := len(a), len(b), len(b[0])
	ret := newMatrix(n, l)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < l; k++ {
				ret[i][k] += a[i][j] * b[j][k]
			}
		}
	}
	return ret
}

func (a matrix) pow(k int) matrix {
	n := len(a)
	if len(a) != len(a[0]) {
		panic(fmt.Errorf("a must be square matrix"))
	}
	ret := newIdentityMatrix(n)
	for k > 0 {
		if k&1 == 1 {
			ret = ret.mul(a)
		}
		a = a.mul(a)
		k >>= 1
	}
	return ret
}
