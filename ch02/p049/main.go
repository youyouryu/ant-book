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
		l0, l1 := popMin(&l), popMin(&l)
		cost += l0 + l1
		l = append(l, l0+l1)
	}
	return
}

func popMin(l *[]int) (min int) {
	slice := *l
	var idx int
	for i := range slice {
		if i == 0 || slice[i] < min {
			min = slice[i]
			idx = i
		}
	}
	*l = append(slice[:idx], slice[idx+1:]...)
	return
}
