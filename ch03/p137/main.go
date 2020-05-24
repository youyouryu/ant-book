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
	a := []int{}
	for i := 0; i < p; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		a = append(a, ai)
	}
	ans := search(a)
	fmt.Fprintln(stdout, ans)
}

func search(a []int) (ans int) {
	p := len(a)
	ans = p
	isCoverd := isCoverdFunc(a)
	first, last := 0, 0
	for first < p {
		for last < p && !isCoverd(first, last) {
			last++
		}
		if !isCoverd(first, last) {
			break
		}
		if last-first+1 < ans {
			ans = last - first + 1
		}
		first++
	}
	return ans
}

func isCoverdFunc(a []int) func(first, last int) bool {
	set := map[int]struct{}{}
	for _, ai := range a {
		set[ai] = struct{}{}
	}
	return func(first, last int) bool {
		for k := range set {
			found := false
			for _, ai := range a[first : last+1] {
				if ai == k {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
		return true
	}
}
