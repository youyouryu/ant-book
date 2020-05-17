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
	ans := isCarmichael(n)
	if ans {
		fmt.Fprintln(stdout, "Yes")
	} else {
		fmt.Fprintln(stdout, "No")
	}
}

func isCarmichael(n int) bool {
	if isPrime(n) {
		return false
	}
	for x := 2; x < n; x++ {
		if pow(x, n, n) != x {
			return false
		}
	}
	return true
}

func isPrime(n int) bool {
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func pow(x, n, mod int) int {
	ans := 1
	for n > 0 {
		if n&1 == 1 {
			ans = ans * x % mod
		}
		x = x * x % mod
		n >>= 1
	}
	return ans
}
