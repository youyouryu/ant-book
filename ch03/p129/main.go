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
	l := []float64{}
	for i := 0; i < n; i++ {
		sc.Scan()
		li, _ := strconv.ParseFloat(sc.Text(), 64)
		l = append(l, li)
	}
	ans := search(n, k, l)
	fmt.Fprintf(stdout, "%.2f\n", ans)
}

func search(n, k int, l []float64) (ans float64) {
	isOk := func(x float64) (ok bool) {
		num := 0
		for i := 0; i < n; i++ {
			num += int(l[i] / x)
		}
		return num >= k
	}

	var sum float64
	for _, li := range l {
		sum += li
	}
	lb, ub := 0.0, sum/float64(k)
	var mid float64
	for i := 0; i < 100; i++ {
		mid = (lb + ub) / 2
		if isOk(mid) {
			lb = mid
		} else {
			ub = mid
		}
	}
	return mid
}
