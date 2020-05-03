package common

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type solverFunc func(r io.Reader, w io.Writer)

// Check the solverFunc can answer correctly
func Check(f solverFunc, qfile string, afile string) error {
	qreader, _ := os.Open(qfile)
	defer qreader.Close()
	areader, _ := os.Open(afile)
	defer areader.Close()
	preader, pwriter := io.Pipe()
	defer preader.Close()
	defer pwriter.Close()
	f(qreader, pwriter)
	return checkEqual(preader, areader)
}

func checkEqual(r1 io.Reader, r2 io.Reader) error {
	sc1 := bufio.NewScanner(r1)
	sc2 := bufio.NewScanner(r2)
	nrows := 0
	for sc1.Scan() {
		if !sc2.Scan() {
			return fmt.Errorf("row number is not equal")
		}
		t1 := sc1.Text()
		t2 := sc2.Text()
		if t1 != t2 {
			return fmt.Errorf("line %d: %s not equals %s", nrows, t1, t2)
		}
		nrows++
	}
	return nil
}
