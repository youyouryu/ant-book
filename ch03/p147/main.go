package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	a, _ := parseArray(sc.Text())
	sc.Scan()
	b, _ := parseArray(sc.Text())
	sc.Scan()
	c, _ := parseArray(sc.Text())
	sc.Scan()
	d, _ := parseArray(sc.Text())
	ans := search(n, a, b, c, d)
	fmt.Println(ans)
}

func parseArray(line string) ([]int, error) {
	arr := []int{}
	values := strings.Split(line, " ")
	for i := range values {
		v, err := strconv.Atoi(values[i])
		if err != nil {
			return arr, err
		}
		arr = append(arr, v)
	}
	return arr, nil
}

func search(n int, a, b, c, d []int) (ans int) {
	// O(n^2)
	e, f := []int{}, []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			e = append(e, a[i]+b[j])
			f = append(f, c[i]+d[j])
		}
	}
	// O(n^2log(n^2)) = O(n^2log(n))
	sort.Slice(e, func(i, j int) bool { return e[i] < e[j] })
	sort.Slice(f, func(i, j int) bool { return f[i] < f[j] })

	// O(n^2log(n^2)) = O(n^2log(n))
	for i := range e {
		lowerBound := sort.Search(len(f), func(j int) bool { return f[j]+e[i] >= 0 })
		upperBound := sort.Search(len(f), func(j int) bool { return f[j]+e[i] > 0 })
		if lowerBound < len(f) {
			ans += upperBound - lowerBound
		}
	}
	return ans
}
