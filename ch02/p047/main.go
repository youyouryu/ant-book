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
	r, _ := strconv.Atoi(sc.Text())
	x := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		xi, _ := strconv.Atoi(sc.Text())
		x = append(x, xi)
	}
	a := greedySearch(n, r, x)
	fmt.Fprintln(stdout, a)
}

func greedySearch(n, r int, x []int) (count int) {
	sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })
	covered := -1
	for i := 0; i < n; i++ {
		if x[i] <= covered {
			continue
		}
		var marked int
		for j := i; j < n; j++ {
			if j == i || x[j] <= x[i]+r {
				marked = x[j]
			}
		}
		covered = marked + r
		count++
	}
	return
}
