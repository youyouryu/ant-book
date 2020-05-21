package common

import (
	"bufio"
	"fmt"
	"io"
	"testing"
)

func TestCheck1(t *testing.T) {
	if err := Check(solve, "test/q1.txt", "test/a1.txt"); err != nil {
		t.Fatal(err)
	}
}

func TestCheck2(t *testing.T) {
	err := Check(solve, "test/q2.txt", "test/a2.txt")
	if err != nil {
		// fmt.Println(err)
	} else {
		t.Fatal("test code could not find errors")
	}
}

func solve(r io.Reader, w io.Writer) {
	sc := bufio.NewScanner(r)
	sc.Scan()
	text := sc.Text()
	fmt.Fprintf(w, "%s %s", text, text)
}
