package common

import (
	"bufio"
	"fmt"
	"io"
	"testing"
)

func TestCheck(t *testing.T) {
	if err := Check(solve, "test/q1.txt", "test/a1.txt"); err != nil {
		t.Fatal(err)
	}
}

func solve(r io.Reader, w io.Writer) {
	sc := bufio.NewScanner(r)
	sc.Scan()
	text := sc.Text()
	fmt.Fprintf(w, "%s %s", text, text)
}
