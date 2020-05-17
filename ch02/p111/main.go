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
	ans := eratosthenes(n)
	fmt.Fprintln(stdout, ans)
}

func eratosthenes(n int) (count int) {
	isPrime := make([]bool, n+1)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false
	for i := 0; i <= n; i++ {
		if !isPrime[i] {
			continue
		}
		for j := 2 * i; j <= n; j += i {
			if j%i == 0 {
				isPrime[j] = false
			}
		}
		count++
	}
	return
}
