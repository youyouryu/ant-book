package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
	a := search(n, wMax, weights, values)
	fmt.Fprintln(stdout, a)
}

func search(n, w int, weights, values []int) int {

	// calculate the upper bound of total value
	vMax := 0
	for _, v := range values {
		vMax += v
	}

	// initialize dp table
	inf := math.MaxInt64
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, vMax+1)
	}
	for j := range dp[0] {
		dp[0][j] = inf
	}
	dp[0][0] = 0

	// run dp
	for i := 0; i < n; i++ {
		for j := 0; j < vMax+1; j++ {
			w0 := dp[i][j]
			if j >= values[i] && dp[i][j-values[i]] < inf {
				w1 := dp[i][j-values[i]] + weights[i]
				if w0 > w1 {
					dp[i+1][j] = w1
					continue
				}
			}
			dp[i+1][j] = w0
		}
	}

	// find best
	for j := vMax; j >= 0; j-- {
		if dp[n][j] <= w {
			return j
		}
	}
	return 0
}
