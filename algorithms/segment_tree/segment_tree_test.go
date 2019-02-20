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

import (
	"fmt"
	"testing"
)

type visitor func(n *node)

func preVisit(n *node, vfun visitor) {
	if n == nil {
		return
	}
	vfun(n)
	preVisit(n.lChild, vfun)
	preVisit(n.rChild, vfun)
}

func TestSegmentTree(t *testing.T) {
	testList := []int{1, 3, 4, 2, 1, 5, 6, 2, 7, 3, 8, 9, 12, 2}
	st := NewSegmentTree(testList)
	str := ""
	vfun := func(n *node) {
		str += fmt.Sprintf("{v: %d, [%d, %d]} ",
			n.value, n.lBound, n.rBound)
	}
	preVisit(st.root, vfun)
	t.Log(str)
	result := st.Query(2, 7)
	if result != 1 {
		t.Errorf("Error Query result, hope 1, but get %d", result)
	}

}
