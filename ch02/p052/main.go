package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(stdin io.Reader, stdout io.Writer) {
	sc := bufio.NewScanner(stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	wMax, _ := strconv.Atoi(sc.Text())
	weights := []int{}
	values := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		line := strings.Split(sc.Text(), " ")
		w, _ := strconv.Atoi(line[0])
		v, _ := strconv.Atoi(line[1])
		weights = append(weights, w)
		values = append(values, v)
	}

	a := search(n, wMax, weights, values, 0)
	fmt.Fprintln(stdout, a)
}

func search(n, w int, weights, values []int, idx int) int {
	if idx == n {
		return 0
	}
	a0 := search(n, w, weights, values, idx+1)
	if w >= weights[idx] {
		a1 := values[idx] + search(n, w-weights[idx], weights, values, idx+1)
		if a0 < a1 {
			return a1
		}
	}
	return a0
}
