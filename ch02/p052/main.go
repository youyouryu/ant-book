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
	memo := map[param]int{}
	a := search(n, wMax, weights, values, 0, &memo)
	fmt.Fprintln(stdout, a)
}

type param struct {
	w   int
	idx int
}

func search(n, w int, weights, values []int, idx int, memo *map[param]int) int {
	if idx == n {
		return 0
	}
	a0 := remindAndMemorize(n, w, weights, values, idx+1, memo)
	if w >= weights[idx] {
		a1 := values[idx] + remindAndMemorize(n, w-weights[idx], weights, values, idx+1, memo)
		if a0 < a1 {
			return a1
		}
	}
	return a0
}

func remindAndMemorize(n, w int, weights, values []int, idx int, memo *map[param]int) (a int) {
	if v, ok := (*memo)[param{w, idx}]; ok {
		a = v
	} else {
		a = search(n, w, weights, values, idx, memo)
		(*memo)[param{w, idx}] = a
	}
	return a
}
