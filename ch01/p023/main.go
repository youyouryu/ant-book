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

func solve(reader io.Reader, writer io.Writer) {
	sc := bufio.NewScanner(reader)
	sc.Scan()
	l, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	x := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		xi, _ := strconv.Atoi(sc.Text())
		x = append(x, xi)
	}
	min, max := calcMinMax(l, x)
	fmt.Fprintln(writer, min)
	fmt.Fprintln(writer, max)
}

func calcMinMax(l int, x []int) (min, max int) {
	mins, maxs := []int{}, []int{}
	for _, xi := range x {
		mins = append(mins, int(math.Min(float64(xi), float64(l-xi))))
		maxs = append(maxs, int(math.Max(float64(xi), float64(l-xi))))
	}
	min = maxInt(mins)
	max = maxInt(maxs)
	return
}

func maxInt(slice []int) (max int) {
	for i, v := range slice {
		if i == 0 || v > max {
			max = v
		}
	}
	return
}
