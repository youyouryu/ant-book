package main

import (
	"fmt"
	"math"
	"os"

	"github.com/youyouryu/AtCoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, c := io.NextInt(), io.NextInt()
	l := io.NextInts(n)
	s := io.NextInts(c)
	a := io.NextInts(c)
	//x, y := solve(l, s, a)
	x, y := solve2(l, s, a)
	for i := range x {
		io.Printf("%.2f %.2f\n", x[i], y[i])
	}

}

// O(n)
func solve(l, s, a []int) (x, y []float64) {
	angles := make([]int, len(l))
	for i := range angles {
		angles[i] = 180
	}
	angles[0] = 90
	for i := range s {
		angles[s[i]] = a[i]
		x[i], y[i] = calcPoint(l, angles)
	}
	return x, y
}

func calcPoint(lengths, angles []int) (x, y float64) {
	sumTheta := 0
	for i := range lengths {
		radius := float64(lengths[i])
		theta := float64((angles[i] - sumTheta + 360) % 360)
		x += radius * math.Cos(2*math.Pi*theta/360)
		y += radius * math.Sin(2*math.Pi*theta/360)
		sumTheta = (sumTheta + angles[i]) % 360
	}
	return x, y
}

// O(log n)
func solve2(l, s, a []int) (x, y []float64) {
	data := make([]interface{}, len(l))
	for i := range data {
		data[i] = node{0, float64(l[i]), 0}
	}
	st := NewSegmentTree(data, merge)
	x, y = make([]float64, len(s)), make([]float64, len(s))
	for i := range s {
		tgt := st.Get(s[i] - 1).(node)
		tgt.theta = float64(a[i])
		st.Update(s[i]-1, tgt)
		top := st.Top().(node)
		x[i], y[i] = top.x, top.y
	}
	return x, y
}

type node struct {
	x, y  float64
	theta float64
}

func merge(a, b interface{}) interface{} {
	na, nb := a.(node), b.(node)
	x := na.x + nb.x*math.Cos(2*math.Pi*na.theta/360) + nb.y*math.Sin(2*math.Pi*na.theta/360)
	y := na.y + nb.x*math.Sin(2*math.Pi*na.theta/360) + nb.y*math.Cos(2*math.Pi*na.theta/360)
	theta := float64(int(na.theta+nb.theta) % 360)
	return node{x, y, theta}
}

type MergeFunc func(interface{}, interface{}) interface{}

type SegmentTree struct {
	offset    int
	nodes     []interface{}
	mergeFunc MergeFunc
}

func NewSegmentTree(data []interface{}, mergeFunc MergeFunc) *SegmentTree {
	depth := 0
	for 1<<depth < len(data) {
		depth++
	}
	size := 1<<(depth+1) - 1
	nodes := make([]interface{}, size)
	offset := 1<<depth - 1
	for i := offset; i < offset+len(data); i++ {
		nodes[i] = data[i-offset]
	}
	s := &SegmentTree{
		offset:    offset,
		nodes:     nodes,
		mergeFunc: mergeFunc,
	}
	return s.fix()
}

func (s *SegmentTree) fix() *SegmentTree {
	for i := s.offset - 1; i >= 0; i-- {
		left, right := s.nodes[2*i+1], s.nodes[2*i+2]
		if right == nil {
			s.nodes[i] = left
			continue
		}
		s.nodes[i] = s.mergeFunc(left, right)
	}
	return s
}

func (s *SegmentTree) Get(index int) interface{} {
	return s.nodes[s.offset+index]
}

func (s *SegmentTree) Top() interface{} {
	return s.nodes[0]
}

func (s *SegmentTree) Update(index int, value interface{}) *SegmentTree {
	s.nodes[s.offset+index] = value
	return s.fix()
}

func (s *SegmentTree) Print() {
	fmt.Printf("offset=%d\n", s.offset)
	for i, v := range s.nodes {
		fmt.Printf("%d %+v\n", i, v)
	}
	fmt.Println()
}
