package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(stdin io.Reader, stdout io.Writer) {
	sc := bufio.NewScanner(stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	k, _ := strconv.Atoi(sc.Text())
	a := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		a = append(a, ai)
	}
	// ans := lowerBound(n, func(i int) bool { return a[i] >= k })
	ans := lowerBound2(n, func(i int) bool { return a[i] >= k })
	fmt.Fprintln(stdout, ans)
}

func lowerBound(n int, f func(int) bool) (idx int) {
	for i := 0; i < n; i++ {
		if f(i) {
			return i
		}
	}
	return n
}

func lowerBound2(n int, f func(int) bool) (idx int) {
	idx = n
	begin, end := 0, n
	for end-begin > 0 {
		pivot := (begin + end) / 2
		if f(pivot) {
			idx = pivot
			end = pivot
		} else {
			begin = pivot + 1
		}
	}
	return idx
}
