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
	arr := [][]rune{}
	for i := 0; i < n; i++ {
		sc.Scan()
		chars := []rune(sc.Text())
		arr = append(arr, chars)
	}
	ans := search(arr)
	fmt.Fprintln(stdout, ans)
}

func search(arr [][]rune) (ans int) {
	n := len(arr)
	fixed := 0
	for j := n - 1; j >= 0; j-- {
		for i := n - 1; i >= 0; i-- {
			if arr[i][j] == '1' && i < j {
				for k := i; k+1 < n-fixed; k++ {
					arr[k], arr[k+1] = arr[k+1], arr[k]
					ans++
				}
				fixed++
			}
		}
	}
	return ans
}
