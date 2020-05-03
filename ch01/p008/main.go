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

func solve(reader io.Reader, writer io.Writer) {
	sc := bufio.NewScanner(reader)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())
	var nums []int
	for i := 0; i < n; i++ {
		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())
		nums = append(nums, num)
	}
	if isWinnable(m, nums, 4) {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}

func isWinnable(m int, nums []int, count int) (ans bool) {
	if count == 0 {
		ans = m == 0
		return
	}
	if m <= 0 {
		return
	}
	for _, k := range nums {
		ans = isWinnable(m-k, nums, count-1)
		if ans {
			return
		}
	}
	return
}
