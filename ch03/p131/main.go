package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	m, _ := strconv.Atoi(sc.Text())
	x := []int{}
	for sc.Scan() {
		xi, _ := strconv.Atoi(sc.Text())
		x = append(x, xi)
	}
	ans := search(n, m, x)
	fmt.Fprintln(stdout, ans)
}

func search(n, m int, x []int) (ans int) {
	sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })
	begin, end := 0, x[n-1]
	var mid int
	for i := 0; i < 100; i++ {
		mid = (begin + end) / 2
		if isOk(x, m, mid) {
			begin = mid
		} else {
			end = mid
		}
	}
	return mid
}

func isOk(x []int, m, d int) (ok bool) {
	cnt := 1
	prev := x[0]
	for i := 1; i < len(x); i++ {
		if x[i] < prev+d {
			continue
		}
		cnt++
		prev = x[i]
	}
	return cnt >= m
}
