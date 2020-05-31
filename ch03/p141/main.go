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
	m, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	stage := make([][]int, m)
	for i := 0; i < m; i++ {
		sc.Scan()
		line := []rune(sc.Text())
		stage[i] = make([]int, n)
		for j := 0; j < n; j++ {
			stage[i][j] = int(line[j] - '0')
		}
	}
	ans := search(stage)
	for i := 0; i < m; i++ {
		line := ""
		for j := 0; j < n; j++ {
			line += fmt.Sprint(ans[i][j])
		}
		fmt.Fprintln(stdout, line)
	}
}

func search(stage [][]int) (ans [][]int) {
	m, n := len(stage), len(stage[0])
	numCasesPerRow := 1 << n
	minCount := m*n + 1

	for k := 0; k < numCasesPerRow; k++ {
		count := 0
		reversed := make([][]int, m)
		for i := 0; i < m; i++ {
			reversed[i] = make([]int, n)
		}
		for j := 0; j < n; j++ {
			if k>>j&1 == 1 {
				reverse(reversed, 0, j)
				count++
			}
		}

		for i := 0; i < m-1; i++ {
			for j := 0; j < n; j++ {
				if (stage[i][j] == 1) == isReversed(reversed, i, j) {
					reverse(reversed, i+1, j)
					count++
				}
			}
		}

		ok := true
		for j := 0; j < n; j++ {
			if (stage[m-1][j] == 1) == isReversed(reversed, m-1, j) {
				ok = false
				break
			}
		}

		if ok && count < minCount {
			minCount = count
			ans = reversed
		}
	}
	return ans
}

func reverse(reversed [][]int, i, j int) {
	m, n := len(reversed), len(reversed[0])
	if i < 0 || m <= i || j < 0 || n <= j {
		return
	}
	reversed[i][j] = (reversed[i][j] + 1) % 2
	return
}

func isReversed(reversed [][]int, i, j int) bool {
	m, n := len(reversed), len(reversed[0])
	num := 0
	if 0 <= i && i < m && 0 <= j && j < n {
		num += reversed[i][j]
	}
	if 0 <= i && i < m && 0 <= j-1 && j-1 < n {
		num += reversed[i][j-1]
	}
	if 0 <= i && i < m && 0 <= j+1 && j+1 < n {
		num += reversed[i][j+1]
	}
	if 0 <= i-1 && i-1 < m && 0 <= j && j < n {
		num += reversed[i-1][j]
	}
	if 0 <= i+1 && i+1 < m && 0 <= j && j < n {
		num += reversed[i+1][j]
	}
	return num%2 == 0
}
