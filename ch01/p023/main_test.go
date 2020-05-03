package main

import (
	"testing"

	"github.com/youyouryu/ant-book/common"
)

func TestSolve1(t *testing.T) {
	if err := common.Check(solve, "test/q1.txt", "test/a1.txt"); err != nil {
		t.Fatal(err)
	}
}
