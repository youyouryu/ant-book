package common

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type solverFunc func(r io.Reader, w io.Writer)

// Check if the solverFunc can answer correctly
func Check(f solverFunc, qfile string, afile string) error {
	qreader, err := os.Open(qfile)
	if err != nil {
		return fmt.Errorf("could not open file: %s", qfile)
	}
	defer qreader.Close()
	areader, err := os.Open(afile)
	if err != nil {
		return fmt.Errorf("could not open file: %s", afile)
	}
	defer areader.Close()
	buffer := new(bytes.Buffer)
	f(qreader, buffer)
	return checkEqual(buffer, areader)
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
