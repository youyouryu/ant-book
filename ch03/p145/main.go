package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

const g = 10.0

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(stdin io.Reader, stdout io.Writer) {
	sc := bufio.NewScanner(stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	h, _ := strconv.ParseFloat(sc.Text(), 64)
	sc.Scan()
	r, _ := strconv.ParseFloat(sc.Text(), 64)
	sc.Scan()
	t, _ := strconv.ParseFloat(sc.Text(), 64)
	ans := calcHeights(n, h, r, t)
	for _, v := range ans {
		fmt.Fprintln(stdout, v)
	}
}

func calcHeights(n int, h, r, t float64) (y []float64) {
	y = []float64{}
	for i := 0; i < n; i++ {
		yi := calcHeight(i, h, r, t)
		y = append(y, yi)
	}
	sort.Slice(y, func(i, j int) bool { return y[i] < y[j] })
	for i := range y {
		y[i] += 2.0 * r * float64(i)
	}
	return y
}

func calcHeight(i int, h, r, t float64) (y float64) {
	fmt.Println(i, h, r, t)
	h += 2.0 * r * float64(i)
	if t < float64(i) {
		return h
	}
	t -= float64(i)
	pt := calcPeriodicTime(h)
	for t > pt {
		t -= pt
	}
	fmt.Println(t, pt)
	if t <= pt/2.0 {
		return h - g*t*t/2.0
	}
	return h - g*(t-pt)*(t-pt)/2.0
}

func calcPeriodicTime(h float64) (periodicTime float64) {
	return math.Sqrt(2.0*h/g) * 2.0
}
