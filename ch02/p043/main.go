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
	s, t := []int{}, []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		line := strings.Split(sc.Text(), " ")
		si, _ := strconv.Atoi(line[0])
		ti, _ := strconv.Atoi(line[1])
		s = append(s, si)
		t = append(t, ti)
	}
	a := greedySearch(n, s, t)
	fmt.Fprintln(stdout, a)
}

func greedySearch(n int, s, t []int) (count int) {
	idxs := []int{}
	for i := 0; i < n; i++ {
		idxs = append(idxs, i)
	}
	sort.Slice(idxs, func(i, j int) bool { return t[i] < t[j] })
	time := 0
	for _, idx := range idxs {
		if s[idx] > time {
			time = t[idx]
			count++
		}
	}
	return
}
