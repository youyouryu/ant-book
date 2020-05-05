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
	l := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		li, _ := strconv.Atoi(sc.Text())
		l = append(l, li)
	}
	a := greedySearch(n, l)
	fmt.Fprintln(stdout, a)
}

func greedySearch(n int, l []int) (cost int) {
	for len(l) > 1 {
		sort.Slice(l, func(i, j int) bool { return l[i] < l[j] })
		cost += l[0] + l[1]
		l = append(l[2:], l[0]+l[1])
	}
	return
}
