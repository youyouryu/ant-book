package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(stdin io.Reader, stdout io.Writer) {
	sc := bufio.NewScanner(stdin)
	sc.Scan()
	s := sc.Text()
	t := bestCowLine(s)
	fmt.Fprintf(stdout, t)
}

func bestCowLine(s string) (t string) {
	sChars := []rune(s)
	tChars := []rune{}
	for i, j := 0, len(sChars)-1; i <= j; {
		if isHeadSmall(sChars[i : j+1]) {
			tChars = append(tChars, sChars[i])
			i++
		} else {
			tChars = append(tChars, sChars[j])
			j--
		}
	}
	t = string(tChars)
	return
}

func isHeadSmall(chars []rune) bool {
	if len(chars) <= 1 {
		return true
	}
	head := 0
	tail := len(chars) - 1
	if chars[head] == chars[tail] {
		return isHeadSmall(chars[head+1 : tail])
	}
	return chars[head] < chars[tail]
}
