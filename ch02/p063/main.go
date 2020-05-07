package main

import (
	"bufio"
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
	a := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		a = append(a, ai)
	}
	ans := search(a)
	fmt.Fprintln(stdout, ans)
}

func search(a []int) (ans int) {
	// 最大増加部分列長 = 系列長 - 重複数 - 反転数
	exists := map[int]bool{}
	ans = len(a)
	for i := 0; i < len(a)-1; i++ {
		if _, ok := exists[a[i]]; ok {
			ans--
		} else {
			exists[a[i]] = true
		}
		if a[i] > a[i+1] {
			ans--
		}
	}
	return
}
