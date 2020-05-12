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
	k, _ := strconv.Atoi(sc.Text())
	r, x, y := []int{}, []int{}, []int{}
	for i := 0; i < k; i++ {
		sc.Scan()
		line := strings.Split(sc.Text(), " ")
		ri, _ := strconv.Atoi(line[0])
		xi, _ := strconv.Atoi(line[1])
		yi, _ := strconv.Atoi(line[2])
		r = append(r, ri)
		x = append(x, xi)
		y = append(y, yi)
	}
	ans := search(n, k, r, x, y)
	fmt.Fprintln(stdout, ans)
}

type unionFind struct {
	par  []int
	rank []int
}

func newUnionFind(n int) unionFind {
	u := unionFind{
		par:  make([]int, n),
		rank: make([]int, n),
	}
	for i := 0; i < n; i++ {
		u.par[i] = i
	}
	return u
}

func (u unionFind) root(x int) int {
	if u.par[x] == x {
		return x
	}
	par := u.root(u.par[x])
	u.par[x] = par
	return par
}

func (u unionFind) unite(x, y int) {
	x = u.root(x)
	y = u.root(y)
	if x == y {
		return
	}

	if u.rank[x] < u.rank[y] {
		u.par[x] = y
	} else {
		u.par[y] = x
		if u.rank[x] == u.rank[y] {
			u.rank[x]++
		}
	}
}

func (u unionFind) same(x, y int) bool {
	return u.root(x) == u.root(y)
}

func search(n, k int, r, x, y []int) (ans int) {
	u := newUnionFind(n * 3)
	for ki := 0; ki < k; ki++ {
		xi := x[ki] - 1
		yi := y[ki] - 1
		ri := r[ki]
		if xi < 0 || n <= xi || yi < 0 || n <= yi {
			ans++
			continue
		}
		if ri == 1 {
			if u.same(xi, yi+n) || u.same(xi, yi+n*2) {
				ans++
			} else {
				u.unite(xi, yi)
				u.unite(xi+n, yi+n)
				u.unite(xi+n*2, yi+n*2)
			}
		} else {
			if u.same(xi, yi) || u.same(xi, yi+n*2) {
				ans++
			} else {
				u.unite(xi, yi+n)
				u.unite(xi+n, yi+n*2)
				u.unite(xi+n*2, yi)
			}
		}
	}
	return
}
