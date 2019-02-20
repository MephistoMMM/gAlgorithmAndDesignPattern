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
package skiplist

import (
	"fmt"
	"testing"
)

func displayList(SL SkipList) {
	sl := SL.(*skipList)
	fmt.Println("\n*****Skip List******")
	head := sl.header
	for lvl := 0; lvl < sl.level+1; lvl++ {
		fmt.Printf("Level %d: \t", lvl)
		node := head.forward[0]
		for node != nil {
			if len(node.forward) > lvl {
				fmt.Printf("%02d\t", node.key)
			} else {

				fmt.Printf("--\t")
			}
			node = node.forward[0]
		}
		fmt.Println("")
	}
}

func TestSkipList(t *testing.T) {
	sl := NewSkipList(10, 0.5)
	sl.Insert(1, 1)
	sl.Insert(10, 1)
	sl.Insert(19, 1)
	sl.Insert(28, 1)
	sl.Insert(90, 1)
	sl.Insert(2, 1)
	sl.Insert(65, 1)
	sl.Insert(23, 1)
	sl.Insert(33, 1)
	sl.Insert(3, 1)
	sl.Insert(17, 1)
	sl.Insert(44, 1)
	displayList(sl)

	sl.Delete(23)
	sl.Delete(33)
	sl.Delete(3)
	displayList(sl)

	if v := sl.Search(65); v == nil {
		t.Errorf("Not found key %d.", 65)
	}

}
