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
	a, _ := strconv.Atoi(line[0])
	b, _ := strconv.Atoi(line[1])
	ans := search(a, b)
	fmt.Fprintln(stdout, intSliceToString(ans))
}

func search(a, b int) (ans []int) {
	var ka, kb int
	if a < b {
		kb, ka = euclid(b, a)
	} else {
		ka, kb = euclid(a, b)
	}
	if ka == 0 && kb == 0 {
		return []int{-1}
	}
	if ka < 0 {
		ans = append(ans, []int{0, -ka}...)
	} else {
		ans = append(ans, []int{ka, 0}...)
	}
	if kb < 0 {
		ans = append(ans, []int{0, -kb}...)
	} else {
		ans = append(ans, []int{kb, 0}...)
	}
	return
}

func euclid(a, b int) (int, int) {
	if b == 0 {
		return 0, 0
	}
	if b == 1 {
		return 0, 1
	}
	q, r := a/b, a%b
	ka, kb := euclid(b, r)
	return kb, ka - kb*q
}

func intSliceToString(slice []int) string {
	str := strconv.Itoa(slice[0])
	for _, v := range slice[1:] {
		str += fmt.Sprintf(" %d", v)
	}
	return str
}
