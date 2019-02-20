// Copyright © 2019 Mephis Pheies <mephistommm@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
package segtree

type node struct {
	value  int
	lBound int
	rBound int
	lChild *node
	rChild *node
}

type SegmentTree struct {
	root *node
	list []int
}

func NewSegmentTree(list []int) *SegmentTree {
	st := &SegmentTree{
		list: list,
	}
	st.init()
	return st
}

func (st *SegmentTree) init() {
	st.root = st.createNode(0, len(st.list))
}

func (st *SegmentTree) calc(lv, rv int) int {
	if lv > rv {
		return rv
	}

	return lv
}

// createNode create a new node
func (st *SegmentTree) createNode(l, r int) *node {
	if l+1 == r {
		return &node{
			value:  st.list[l],
			lBound: l,
			rBound: r,
		}
	}

	mid := (r + l) / 2
	lchild := st.createNode(l, mid)
	rchild := st.createNode(mid, r)
	newNode := &node{
		value:  st.calc(lchild.value, rchild.value),
		lBound: l,
		rBound: r,
		lChild: lchild,
		rChild: rchild,
	}

	return newNode
}

func (st *SegmentTree) Query(l, r int) int {
	return st.queryResultIn(st.root, l, r)
}

func (st *SegmentTree) queryResultIn(n *node, x, y int) int {
	if x <= n.lBound && n.rBound <= y {
		return n.value
	}

	mid := (n.lBound + n.rBound) / 2
	// l  x    y  mid
	// |--|....|--|
	if y <= mid {
		return st.queryResultIn(n.lChild, x, y)
	}
	// mid  x    y   r
	// |----|....|---|
	if mid <= x {
		return st.queryResultIn(n.rChild, x, y)
	}

	// l   x   mid  y  r
	// |---|........|--|
	return st.calc(
		st.queryResultIn(n.lChild, x, mid),
		st.queryResultIn(n.rChild, mid, y))
}
