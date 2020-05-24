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
	s, _ := strconv.Atoi(sc.Text())
	a := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		a = append(a, ai)
	}
	// ans := search(s, a)
	// ans := search2(s, a)
	ans := search3(s, a)
	fmt.Fprintln(stdout, ans)
}

func search(s int, a []int) (ans int) {
	n := len(a)
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j+i >= n {
				continue
			}
			dp[j] = dp[j+1] + a[j]
			if dp[j] >= s {
				return i + 1
			}
		}
	}
	return -1
}

func search2(s int, a []int) (ans int) {
	n := len(a)
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + a[i]
	}
	if sum[n] < s {
		return -1
	}
	min := n
	for i := 0; i < n; i++ {
		j := sort.Search(n, func(j int) bool { return sum[j]-sum[i] >= s })
		if j-i < min {
			min = j - i + 1
		}
	}
	return min
}

func search3(s int, a []int) (ans int) {
	n := len(a)
	ans = n + 1
	first, last, sum := 0, 0, 0
	for {
		for last < n && sum < s {
			sum += a[last]
			last++
		}
		if sum < s {
			break
		}
		if last-first < ans {
			ans = last - first
		}
		sum -= a[first]
		first++
	}
	if ans > n {
		return -1
	}
	return ans
}
