package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(stdin io.Reader, stdout io.Writer) {
	sc := bufio.NewScanner(stdin)
	sc.Scan()
	line := strings.Split(sc.Text(), " ")
	a, _ := strconv.Atoi(line[0])
	b, _ := strconv.Atoi(line[1])
	ans := eratosthenes(a, b)
	fmt.Fprintln(stdout, ans)
}

func eratosthenes(a, b int) (count int) {
	limit := int(math.Sqrt(float64(b)))
	isPrimeSmall := make([]bool, limit)
	for i := range isPrimeSmall {
		isPrimeSmall[i] = true
	}
	isPrimeSmall[0], isPrimeSmall[1] = false, false

	isPrime := make([]bool, b-a)
	for i := range isPrime {
		isPrime[i] = true
	}

	for i := 2; i < limit; i++ {
		if !isPrimeSmall[i] {
			continue
		}
		for j := 2 * i; j < limit; j += i {
			if j%i == 0 {
				isPrimeSmall[j] = false
			}
		}
		begin := (a+i-1)/i*i - a // a以上の最小のiの倍数
		for j := begin; j < b-a; j += i {
			isPrime[j] = false
		}
	}
	for i := 0; i < b-a; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}
