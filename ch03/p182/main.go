package main

import (
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
	ans := solver(n)
	io.Println(ans)
}

func solver(n int) int {
	a := lib.Matrix([][]int{{2, 1, 0}, {2, 2, 1}, {0, 1, 2}})
	b := lib.Matrix([][]int{{2}, {2}, {0}})
	return a.Pow(n - 1).Mul(b)[0][0]
}
