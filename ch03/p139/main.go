package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(stdin io.Reader, stdout io.Writer) {
	sc := bufio.NewScanner(stdin)
	sc.Scan()
	cows := []rune(sc.Text())
	k, m := search(cows)
	fmt.Fprintln(stdout, k, m)
}

func search(cows []rune) (kMin, mMin int) {
	kMin, mMin = math.MaxInt64, math.MaxInt64
	n := len(cows)
	for k := 0; k <= n; k++ {
		cp := make([]rune, n)
		copy(cp, cows)
		m := 0
		found := true
		for i := 0; i < n; i++ {
			if cp[i] == 'F' {
				continue
			}
			if k == 0 || i+k > n {
				found = false
				break
			}
			toggleRange(cp, i, i+k)
			m++
		}
		if found && m < mMin {
			kMin, mMin = k, m
		}
	}
	return kMin, mMin
}

func toggle(cows []rune, i int) {
	if cows[i] == 'F' {
		cows[i] = 'B'
	} else if cows[i] == 'B' {
		cows[i] = 'F'
	}
}

func toggleRange(cows []rune, begin, end int) {
	for i := begin; i < end; i++ {
		toggle(cows, i)
	}
}
