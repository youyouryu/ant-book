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
	v1, v2 := []int{}, []int{}
	sc.Scan()
	line1 := strings.Split(sc.Text(), " ")
	sc.Scan()
	line2 := strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		v1i, _ := strconv.Atoi(line1[i])
		v1 = append(v1, v1i)
		v2i, _ := strconv.Atoi(line2[i])
		v2 = append(v2, v2i)
	}
	ans := minimumDotProduct(v1, v2)
	fmt.Fprintln(stdout, ans)
}

func minimumDotProduct(v1, v2 []int) (ans int) {
	sort.Slice(v1, func(i, j int) bool { return v1[i] < v1[j] })
	sort.Slice(v2, func(i, j int) bool { return v2[i] > v2[j] })
	for i := 0; i < len(v1); i++ {
		ans += v1[i] * v2[i]
	}
	return ans
}
