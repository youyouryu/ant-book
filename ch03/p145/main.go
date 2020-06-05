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
		fmt.Fprintf(stdout, "%.2f\n", v)
	}
}

func calcHeights(n int, h, r, t float64) (y []float64) {
	y = []float64{}
	for i := 0; i < n; i++ {
		//yi := calcHeight(h+2*r*float64(i), t-float64(i))
		yi := calcHeight(h, t-float64(i))
		y = append(y, yi)
	}
	sort.Slice(y, func(i, j int) bool { return y[i] < y[j] })
	for i := range y {
		y[i] += 2.0 * r * float64(i)
	}
	return y
}

func calcHeight(h, t float64) (y float64) {
	if t < 0 {
		return h
	}
	ft := calcFallTime(h)
	k := int(t / ft)
	var d float64
	if k%2 == 0 {
		d = t - float64(k)*ft
	} else {
		d = float64(k)*ft + ft - t
	}
	return h - g*d*d*0.5
}

func calcFallTime(h float64) (periodicTime float64) {
	return math.Sqrt(2.0 * h / g)
}
