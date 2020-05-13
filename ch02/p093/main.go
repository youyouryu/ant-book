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
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		sc.Scan()
		line := strings.Split(sc.Text(), " ")
		vFrom, _ := strconv.Atoi(line[0])
		for j := range line[1:] {
			vTo, _ := strconv.Atoi(line[1:][j])
			graph[vFrom][vTo] = 1
		}
	}
	colors := make([]int, n)
	ans := isPaintable(&graph, colors, 0, 1)
	if ans {
		fmt.Fprintln(stdout, "Yes")
	} else {
		fmt.Fprintln(stdout, "No")
	}
}

func isPaintable(graph *[][]int, colors []int, vertex, color int) bool {
	colors[vertex] = color
	for v, e := range (*graph)[vertex] {
		if e == 0 {
			continue
		}
		if colors[v] == color {
			return false
		}
		if colors[v] == 0 && !isPaintable(graph, colors, v, -color) {
			return false
		}
	}
	return true
}
