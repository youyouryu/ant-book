package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
	ans := search(n, m, rows)
	fmt.Fprintln(stdout, ans)
}

func search(n, m int, rows [][]rune) int {
	moves := []point{
		{x: -1, y: 0},
		{x: 1, y: 0},
		{x: 0, y: -1},
		{x: 0, y: 1},
	}
	distances := initDistances(n, m)
	queue := []point{}
	p := findRune(n, m, rows, 'S')
	distances[p.x][p.y] = 0
	queue = append(queue, p)

	for len(queue) > 0 {
		p = queue[0]
		queue = queue[1:]
		if rows[p.x][p.y] == 'G' {
			return distances[p.x][p.y]
		}
		for _, dp := range moves {
			np := point{x: p.x + dp.x, y: p.y + dp.y}
			if np.x < 0 || n <= np.x || np.y < 0 || m <= np.y {
				continue
			}
			if rows[np.x][np.y] == '#' {
				continue
			}
			if distances[np.x][np.y] != math.MaxInt64 {
				continue
			}
			distances[np.x][np.y] = distances[p.x][p.y] + 1
			queue = append(queue, np)
		}
	}
	return -1
}

type point struct {
	x int
	y int
}

func initDistances(n, m int) [][]int {
	rows := [][]int{}
	for i := 0; i < n; i++ {
		row := []int{}
		for j := 0; j < m; j++ {
			row = append(row, math.MaxInt64)
		}
		rows = append(rows, row)
	}
	return rows
}

func findRune(n, m int, rows [][]rune, r rune) point {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if rows[i][j] == r {
				return point{x: i, y: j}
			}
		}
	}
	return point{x: -1, y: -1}
}
