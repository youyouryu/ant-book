package main

import (
	"testing"

	"github.com/yuyamada/ant-book/common"
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
