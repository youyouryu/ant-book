package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
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
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func isWinnable(m int, nums []int, count int) (ans bool) {
	fmt.Println(m, nums, count)
	if count == 0 {
		ans = m == 0
		return
	}
	if m <= 0 {
		return
	}
	for _, k := range nums {
		ans = ans || isWinnable(m-k, nums, count-1)
		if ans {
			return
		}
	}
	return
}
