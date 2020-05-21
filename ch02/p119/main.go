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
	pos := []int{}
	n := len(arr)
	for i := 0; i < n; i++ {
		p := -1
		for j := 0; j < n; j++ {
			if arr[i][j] == '1' {
				p = j
			}
		}
		pos = append(pos, p)
	}

	for i := 0; i < n; i++ {
		var j int
		for j = i; j < n; j++ {
			if pos[j] <= i {
				break
			}
		}
		for k := j; k-1 >= i; k-- {
			pos[k-1], pos[k] = pos[k], pos[k-1]
			ans++
		}
	}
	return ans
}
