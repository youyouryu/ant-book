package main

import (
	"bufio"
	"container/heap"
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

type edge struct {
	to   int
	cost int
}

type vertex struct {
	id   int
	dist int
}

type priorityQueue []vertex

func (p priorityQueue) Len() int               { return len(p) }
func (p priorityQueue) Less(i int, j int) bool { return p[i].dist < p[j].dist }
func (p priorityQueue) Swap(i, j int)          { p[i], p[j] = p[j], p[i] }
func (p *priorityQueue) Push(x interface{})    { *p = append(*p, x.(vertex)) }
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
	r, _ := strconv.Atoi(sc.Text())
	graph := make([][]edge, n)
	for i := range graph {
		graph[i] = []edge{}
	}
	for i := 0; i < r; i++ {
		sc.Scan()
		line := strings.Split(sc.Text(), " ")
		from, _ := strconv.Atoi(line[0])
		to, _ := strconv.Atoi(line[1])
		cost, _ := strconv.Atoi(line[2])
		graph[from-1] = append(graph[from-1], edge{to: to - 1, cost: cost})
		graph[to-1] = append(graph[to-1], edge{to: from - 1, cost: cost})
	}
	ans := dijkstra(graph)
	fmt.Fprintln(stdout, ans)
}

func dijkstra(graph [][]edge) int {
	dist := make([]int, len(graph))
	dist2 := make([]int, len(graph))
	for i := range graph {
		if i == 0 {
			dist[i], dist2[i] = 0, math.MaxInt64
			continue
		}
		dist[i], dist2[i] = math.MaxInt64, math.MaxInt64
	}
	queue := priorityQueue{}
	heap.Push(&queue, vertex{id: 0, dist: 0})
	for len(queue) > 0 {
		v := heap.Pop(&queue).(vertex)
		if dist2[v.id] <= v.dist {
			continue
		}
		for j := 0; j < len(graph[v.id]); j++ {
			e := graph[v.id][j]
			d := dist[v.id] + e.cost
			if d < dist[e.to] {
				dist2[e.to] = dist[e.to]
				dist[e.to] = d
				heap.Push(&queue, vertex{id: e.to, dist: dist[e.to]})
				heap.Push(&queue, vertex{id: e.to, dist: dist2[e.to]})
			} else if d < dist2[e.to] {
				dist2[e.to] = d
				heap.Push(&queue, vertex{id: e.to, dist: dist2[e.to]})
			}
		}
	}
	return dist2[len(dist2)-1]
}
