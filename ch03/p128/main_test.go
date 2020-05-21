package main

import (
	"testing"

	"github.com/youyouryu/ant-book/common"
)

func TestSolve1(t *testing.T) {
	common.Check(solve, "test/1.in", "test/1.out")
}

func TestSolve2(t *testing.T) {
	common.Check(solve, "test/2.in", "test/2.out")
}

func TestSolve3(t *testing.T) {
	common.Check(solve, "test/3.in", "test/3.out")
}

func TestSolve4(t *testing.T) {
	common.Check(solve, "test/4.in", "test/4.out")
}
