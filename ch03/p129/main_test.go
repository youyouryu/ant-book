package main

import (
	"testing"

	"github.com/youyouryu/ant-book/common"
)

func TestSolve1(t *testing.T) {
	if err := common.Check(solve, "test/1.in", "test/1.out"); err != nil {
		t.Error(err)
	}
}
