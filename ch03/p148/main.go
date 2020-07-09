package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(stdin io.Reader, stdout io.Writer) {
	sc := bufio.NewScanner(stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())
	w, v := []int{}, []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		line := strings.Split(sc.Text(), " ")
		wi, _ := strconv.Atoi(line[0])
		vi, _ := strconv.Atoi(line[1])
		w, v = append(w, wi), append(v, vi)
	}
	ans := search(n, m, w, v)
	fmt.Fprintln(stdout, ans)
}

func search(n, m int, w, v []int) (ans int) {
	w1, v1 := combinations(w[:n/2], v[:n/2])
	w2, v2 := combinations(w[n/2:], v[n/2:])
	indices1 := sortIndices(w1, v1)
	indices2 := sortIndices(w2, v2)
	max := 0
	for _, idx1 := range indices1 {
		i := sort.Search(len(indices2), func(i int) bool {
			idx2 := indices2[i]
			return w1[idx1]+w2[idx2] > m
		})
		idx2 := indices2[i-1]
		if v1[idx1]+v2[idx2] > max {
			max = v1[idx1] + v2[idx2]
		}
	}
	return max
}

func combinations(w, v []int) (wNew, vNew []int) {
	for i, maxIter := 0, 1<<len(w); i < maxIter; i++ {
		wi, vi := 0, 0
		for j := 0; j < len(w); j++ {
			if (i>>j)&1 == 1 {
				wi += w[j]
				vi += v[j]
			}
		}
		wNew = append(wNew, wi)
		vNew = append(vNew, vi)
	}
	return wNew, vNew
}

func sortIndices(w, v []int) (indices []int) {
	indices = make([]int, len(w))
	for i := range indices {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		if w[i] < w[j] {
			return true
		} else if w[i] == w[j] {
			return v[i] < v[j]
		} else {
			return false
		}
	})
	return indices
}
