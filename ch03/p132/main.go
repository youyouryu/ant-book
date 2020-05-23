package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	k, _ := strconv.Atoi(sc.Text())
	w, v := []int{}, []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		line := strings.Split(sc.Text(), " ")
		wi, _ := strconv.Atoi(line[0])
		vi, _ := strconv.Atoi(line[1])
		w, v = append(w, wi), append(v, vi)
	}
	ans := search(n, k, w, v)
	fmt.Fprintln(stdout, ans)
}

func search(n, k int, w, v []int) (ans float64) {
	indices := []int{}
	for i := 0; i < n; i++ {
		indices = append(indices, i)
	}
	isOk := func(x float64) bool {
		sort.Slice(indices, func(i, j int) bool {
			return float64(v[i])-x*float64(w[i]) > float64(v[j])-x*float64(w[j])
		})
		sum := 0.0
		for i := 0; i < k; i++ {
			j := indices[i]
			sum += float64(v[j]) - x*float64(w[j])
		}
		return sum >= 0
	}
	s := 0.0
	for _, vi := range v {
		s += float64(vi)
	}
	begin, end := 0.0, s
	var mid float64
	for i := 0; i < 100; i++ {
		mid = (begin + end) / 2
		if isOk(mid) {
			begin = mid
		} else {
			end = mid
		}
	}
	return mid
}
