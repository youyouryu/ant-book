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
	yens := []int{1, 5, 10, 50, 100, 500}
	nums := []int{}
	for i := 0; i < len(yens); i++ {
		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())
		nums = append(nums, num)
	}
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	cnt := greedySearch(yens, nums, a)
	fmt.Fprintln(stdout, cnt)
}

func greedySearch(yens, nums []int, a int) (count int) {
	for i := len(yens) - 1; i >= 0; i-- {
		for j := 0; j < nums[i]; j++ {
			if yens[i] > a {
				break
			}
			a -= yens[i]
			count++
		}
	}
	return
}
