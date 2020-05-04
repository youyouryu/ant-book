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
	sc.Scan()
	k, _ := strconv.Atoi(sc.Text())
	nums := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())
		nums = append(nums, num)
	}
	if search(nums, k, 0) {
		fmt.Fprintln(stdout, "Yes")
	} else {
		fmt.Fprintln(stdout, "No")
	}
}

func search(nums []int, k int, idx int) bool {
	if idx == len(nums) {
		return k == 0
	}
	if search(nums, k-nums[idx], idx+1) || search(nums, k, idx+1) {
		return true
	}
	return false
}
