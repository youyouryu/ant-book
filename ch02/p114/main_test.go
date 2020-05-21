package main

import (
	"testing"

	"github.com/youyouryu/ant-book/common"
)

func TestSolve1(t *testing.T) {
	if err := common.Check(solve, "test/q1.txt", "test/a1.txt"); err != nil {
		t.Error(err)
	}
}

func TestSolve2(t *testing.T) {
	if err := common.Check(solve, "test/q2.txt", "test/a2.txt"); err != nil {
		t.Error(err)
	}
}

func TestSolve3(t *testing.T) {
	if err := common.Check(solve, "test/q3.txt", "test/a3.txt"); err != nil {
		t.Error(err)
	}
}
