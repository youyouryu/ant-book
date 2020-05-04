package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	if search(nums, k) {
		fmt.Fprintln(stdout, "Yes")
	} else {
		fmt.Fprintln(stdout, "No")
	}
}

func search(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}
	if k == 0 {
		return true
	} else if k < 0 {
		return false
	}
	for i := range nums {
		if search(append(nums[:i], nums[i+1:]...), k-nums[i]) {
			return true
		}
	}
	return false
}
