package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	edges := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		e, _ := strconv.Atoi(sc.Text())
		edges = append(edges, e)
	}
	fmt.Println(biggestCircumference(edges))
}

func biggestCircumference(edges []int) (bc int) {
	sort.Slice(edges, func(i, j int) bool { return edges[i] > edges[j] })
	for i := 0; i < len(edges)-2; i++ {
		if isTriangle(edges[i], edges[i+1], edges[i+2]) {
			bc = edges[i] + edges[i+1] + edges[i+2]
			return
		}
	}
	return
}

func isTriangle(a, b, c int) (ans bool) {
	slice := []int{a, b, c}
	sort.Slice(slice, func(i, j int) bool { return slice[i] > slice[j] })
	ans = slice[0] < slice[1]+slice[2]
	return
}
