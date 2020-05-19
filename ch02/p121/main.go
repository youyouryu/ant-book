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
	p, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())
	a := []int{}
	for sc.Scan() {
		ai, _ := strconv.Atoi(sc.Text())
		a = append(a, ai)
	}
	ans := calcCost(p, q, a)
	fmt.Fprintln(stdout, ans)
}

func calcCost(p, q int, a []int) (cost int) {
	queue := []int{}
	a = append(a, p+1)
	prev := 0
	for i := range a {
		v := a[i] - prev - 1
		queue = append(queue, v)
		prev = a[i]
	}
	for len(queue) > 1 {
		var min, argmin int
		for i := 0; i < len(queue)-1; i++ {
			c := queue[i] + queue[i+1]
			if i == 0 || c < min {
				min, argmin = c, i
			}
		}
		cost += min
		tmp := append(queue[:argmin], min+1)
		queue = append(tmp, queue[argmin+2:]...)
	}
	return cost
}
