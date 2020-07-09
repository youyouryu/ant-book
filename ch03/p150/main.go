package main

import (
	"os"
	"sort"

	"github.com/yuyamada/AtCoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	w, h, n := io.NextInt(), io.NextInt(), io.NextInt()
	x1 := io.NextInts(n)
	x2 := io.NextInts(n)
	y1 := io.NextInts(n)
	y2 := io.NextInts(n)
	ans := solve(w, h, x1, x2, y1, y2)
	io.Println(ans)
}

func solve(w, h int, x1, x2, y1, y2 []int) (ans int) {
	stage := createStage(w, h, x1, x2, y1, y2)
	for i := range stage {
		for j := range stage[i] {
			if stage[i][j] == 1 {
				continue
			}
			fill(&stage, i, j)
			ans++
		}
	}
	return ans
}

func createStage(w, h int, x1, x2, y1, y2 []int) [][]int {
	mx := compress(append(x1, x2...), 1, w)
	my := compress(append(y1, y2...), 1, h)
	nRows, nCols := len(mx), len(my)
	stage := make([][]int, nRows)
	for i := 0; i < nRows; i++ {
		stage[i] = make([]int, nCols)
	}
	for i := range x1 {
		cx1, cy1 := mx[x1[i]], my[y1[i]]
		cx2, cy2 := mx[x2[i]], my[y2[i]]
		if cx1 > cx2 {
			cx1, cx2 = cx2, cx1
		}
		if cy1 > cy2 {
			cy1, cy2 = cy2, cy1
		}
		for j := cx1; j <= cx2; j++ {
			for k := cy1; k <= cy2; k++ {
				stage[j][k] = 1
			}
		}
	}
	return stage
}

func compress(values []int, min, max int) map[int]int {
	mp := map[int]int{}
	sort.Ints(values)
	for _, v := range values {
		for dv := -1; dv <= 1; dv++ {
			if v+dv < min || max < v+dv {
				continue
			}
			if _, ok := mp[v+dv]; ok {
				continue
			}
			mp[v+dv] = len(mp)
		}
	}
	return mp
}

var dx []int = []int{-1, 1, 0, 0}
var dy []int = []int{0, 0, -1, 1}

func fill(stage *[][]int, x, y int) {
	w, h := len(*stage), len((*stage)[0])
	(*stage)[x][y] = 1
	for i := range dx {
		if x+dx[i] < 0 || w <= x+dx[i] || y+dy[i] < 0 || h <= y+dy[i] {
			continue
		}
		if (*stage)[x+dx[i]][y+dy[i]] == 1 {
			continue
		}
		fill(stage, x+dx[i], y+dy[i])
	}
}
