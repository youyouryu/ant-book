package main

import (
	"bufio"
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

func solve(stdin io.Reader, stdout io.Writer) {
	sc := bufio.NewScanner(stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	ml, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	md, _ := strconv.Atoi(sc.Text())

	al, bl, dl := []int{}, []int{}, []int{}
	for i := 0; i < ml; i++ {
		sc.Scan()
		line := strings.Split(sc.Text(), " ")
		ali, _ := strconv.Atoi(line[0])
		bli, _ := strconv.Atoi(line[1])
		dli, _ := strconv.Atoi(line[2])
		al = append(al, ali)
		bl = append(bl, bli)
		dl = append(dl, dli)
	}
	ad, bd, dd := []int{}, []int{}, []int{}
	for i := 0; i < md; i++ {
		sc.Scan()
		line := strings.Split(sc.Text(), " ")
		adi, _ := strconv.Atoi(line[0])
		bdi, _ := strconv.Atoi(line[1])
		ddi, _ := strconv.Atoi(line[2])
		ad = append(ad, adi)
		bd = append(bd, bdi)
		dd = append(dd, ddi)
	}

	a := search(n, ml, md, al, bl, dl, ad, bd, dd)
	fmt.Fprintln(stdout, a)
}

func search(n, ml, md int, al, bl, dl, ad, bd, dd []int) int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[0] = 0
	for k := 0; k < n; k++ {
		for i := 0; i+1 < n; i++ {
			if dist[i+1] == math.MaxInt64 {
				continue
			}
			if dist[i+1] < dist[i] {
				dist[i] = dist[i+1]
			}
		}
		for i := 0; i < ml; i++ {
			if dist[al[i]-1] == math.MaxInt64 {
				continue
			}
			d := dist[al[i]-1] + dl[i]
			if d < dist[bl[i]-1] {
				dist[bl[i]-1] = d
			}
		}
		for i := 0; i < md; i++ {
			if dist[bd[i]-1] == math.MaxInt64 {
				continue
			}
			d := dist[bd[i]-1] - dd[i]
			if d < dist[ad[i]-1] {
				dist[ad[i]-1] = d
			}
		}
	}
	if dist[n-1] < 0 {
		return -1
	} else if dist[n-1] == math.MaxInt64 {
		return -2
	}
	return dist[n-1]
}
