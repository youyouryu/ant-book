package main

import (
	"bufio"
	"fmt"
	"io"
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
	x1, _ := strconv.Atoi(line[0])
	y1, _ := strconv.Atoi(line[1])
	sc.Scan()
	line = strings.Split(sc.Text(), " ")
	x2, _ := strconv.Atoi(line[0])
	y2, _ := strconv.Atoi(line[1])
	ans := search(x1, y1, x2, y2)
	fmt.Fprintln(stdout, ans)
}

func search(x1, y1, x2, y2 int) int {
	a := abs(x2 - x1)
	b := abs(y2 - y1)
	if a == 0 && b == 0 {
		return 0
	}
	if a < b {
		return gcd(b, a) - 1
	}
	return gcd(a, b) - 1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
