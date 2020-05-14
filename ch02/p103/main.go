package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

type edge struct {
	to   int
	cost int
}

type priorityQueue []edge

func (p priorityQueue) Len() int            { return len(p) }
func (p priorityQueue) Less(i, j int) bool  { return p[i].cost < p[j].cost }
func (p priorityQueue) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *priorityQueue) Push(x interface{}) { *p = append(*p, x.(edge)) }
func (p *priorityQueue) Pop() interface{} {
	old := *p
	*p = old[:len(old)-1]
	return old[len(old)-1]
}

func solve(stdin io.Reader, stdout io.Writer) {
	sc := bufio.NewScanner(stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	graph := make([][]int, n+m)
	for i := range graph {
		graph[i] = make([]int, n+m)
	}
	for i := 0; i < r; i++ {
		sc.Scan()
		line := strings.Split(sc.Text(), " ")
		from, _ := strconv.Atoi(line[0])
		to, _ := strconv.Atoi(line[1])
		closeness, _ := strconv.Atoi(line[2])
		if graph[from][to+n] < -closeness {
			continue
		}
		graph[from][to+n] = -closeness
		graph[to+n][from] = -closeness
	}

	ans := search(&graph)
	fmt.Fprintln(stdout, ans)
}

func search(graph *[][]int) (ans int) {
	g := *graph
	contains := make([]bool, len(g))
	queue := priorityQueue{}

	heap.Push(&queue, edge{to: 0, cost: 0})
	for len(queue) > 0 {
		e := heap.Pop(&queue).(edge)
		if contains[e.to] {
			continue
		}
		ans += 10000 + e.cost
		contains[e.to] = true
		for j := range g[e.to] {
			if contains[j] {
				continue
			}
			heap.Push(&queue, edge{to: j, cost: g[e.to][j]})
		}
	}
	return
}
