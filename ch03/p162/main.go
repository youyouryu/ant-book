package main

import (
	"fmt"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	a := io.NextInts(n)
	ans := solve(a)
	io.Println(ans)
}

func solve(a []int) (ans int) {
	fmt.Println(a)
	bit := lib.NewBinaryIndexedTree(len(a))
	for i := range a {
		ans += i - bit.Sum(a[i])
		bit.Add(a[i], 1)

		fmt.Println(bit)
	}
	return ans
}
