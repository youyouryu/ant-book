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
	a := io.NextInts(n)
	ans := solver(a)
	io.Println(ans)
}

func solver(a []int) (ans int) {
	bit := lib.NewBinaryIndexedTree(len(a))
	for i := range a {
		ans += i - bit.Sum(a[i])
		bit.Add(a[i], 1)
	}
	return ans
}
