package main

import (
	"testing"

	"github.com/yuyamada/ant-book/common"
)

func TestSolve1(t *testing.T) {
	if err := common.Check(solve, "test/1.in", "test/1.out"); err != nil {
		t.Error(err)
	}
}

func TestSolve2(t *testing.T) {
	if err := common.Check(solve, "test/2.in", "test/2.out"); err != nil {
		t.Error(err)
	}
}
