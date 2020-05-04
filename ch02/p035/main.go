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
	m, _ := strconv.Atoi(sc.Text())
	rows := [][]rune{}
	for sc.Scan() {
		row := []rune(sc.Text())
		rows = append(rows, row)
	}
	ans := search(n, m, &rows)
	fmt.Fprintln(stdout, ans)
}

func search(n, m int, rows *[][]rune) (count int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (*rows)[i][j] == 'W' {
				dfs(i, j, n, m, rows)
				count++
			}
		}
	}
	return
}

func dfs(i, j, n, m int, rows *[][]rune) {
	(*rows)[i][j] = '.'
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}
			if i+di < 0 || n <= i+di || j+dj < 0 || m <= j+dj {
				continue
			}
			if (*rows)[i+di][j+dj] == 'W' {
				dfs(i+di, j+dj, n, m, rows)
			}
		}
	}
}
