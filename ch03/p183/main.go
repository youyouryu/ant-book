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
	n, m := io.NextInt(), io.NextInt()
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		from, to := io.NextInt()-1, io.NextInt()-1
		g[from][to] = 1
	}
	k := io.NextInt()
	ans := solver(g, k)
	io.Println(ans)
}

func solver(g [][]int, k int) (ans int) {
	mat := lib.Matrix(g)
	mat = mat.Pow(k)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			ans += mat[i][j]
		}
	}
	return ans
}
