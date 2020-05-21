package main

import (
	"bufio"
	"container/heap"
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
	l := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		li, _ := strconv.Atoi(sc.Text())
		l = append(l, li)
	}
	// a := greedySearch(n, l)
	a := searchWithHeap(n, l)
	fmt.Fprintln(stdout, a)
}

func greedySearch(n int, l []int) (cost int) {
	for len(l) > 1 {
		l0, l1 := popMin(&l), popMin(&l)
		cost += l0 + l1
		l = append(l, l0+l1)
	}
	return
}

func popMin(l *[]int) (min int) {
	slice := *l
	var idx int
	for i := range slice {
		if i == 0 || slice[i] < min {
			min = slice[i]
			idx = i
		}
	}
	*l = append(slice[:idx], slice[idx+1:]...)
	return
}

type intHeap []int

func (h intHeap) Len() int            { return len(h) }
func (h intHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func searchWithHeap(n int, l []int) (cost int) {
	h := intHeap(l)
	for h.Len() > 1 {
		l0, l1 := heap.Pop(&h).(int), heap.Pop(&h).(int)
		heap.Push(&h, l0+l1)
		cost += l0 + l1
	}
	return cost
}
