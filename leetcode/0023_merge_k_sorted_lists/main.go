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
package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/merge-k-sorted-lists/

// https://leetcode.com/problems/merge-k-sorted-lists/solution/
// Approach 3: Optimize Approach 2 by Priority Queue

type ListNode struct {
	Val  int
	Next *ListNode
}

func compare(a, b *ListNode) int {
	if a == nil {
		return 1
	}

	if b == nil {
		return -1
	}

	return a.Val - b.Val
}

type loserTree struct {
	tree  []int
	lists []*ListNode
	size  int
}

func NewLoserTree(lists []*ListNode) *loserTree {
	size := len(lists)
	lt := &loserTree{
		size:  size,
		tree:  make([]int, size), // size-1 + 1
		lists: make([]*ListNode, size),
	}
	copy(lt.lists, lists)

	lt.init()
	return lt
}

func (lt *loserTree) getNode(i int) *ListNode {
	if i >= lt.size {
		return lt.lists[i-lt.size]
	}

	return lt.lists[lt.tree[i]-lt.size]
}

func (lt *loserTree) init() {

	// init loser tree
	for i := 2*lt.size - 1; i >= lt.size; i-- {
		cur := i
		for j := cur; j > 0; j /= 2 {
			parent := j / 2
			if lt.tree[parent] == 0 {
				lt.tree[parent] = cur
				break
			}

			if compare(lt.getNode(parent), lt.getNode(cur)) < 0 {
				tmp := cur
				cur = lt.tree[parent]
				lt.tree[parent] = tmp
			}
		}
	}
}

// contest ...
func (lt *loserTree) contest() (int, bool) {
	if lt.getNode(0) == nil {
		return 0, false
	}

	cur := lt.tree[0]
	result := lt.getNode(0).Val
	lt.lists[cur-lt.size] = lt.lists[cur-lt.size].Next
	for i := cur; i > 0; i /= 2 {
		parent := i / 2
		if compare(lt.getNode(parent), lt.getNode(cur)) < 0 {
			tmp := cur
			cur = lt.tree[parent]
			lt.tree[parent] = tmp
		}
	}

	if cur != lt.tree[0] {
		lt.tree[0] = cur
	}

	return result, true
}

// 12 ms 6.1 MB
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	rhead := &ListNode{}
	rnode := rhead
	lt := NewLoserTree(lists)
	fmt.Printf("%v \n ", lt.tree)

	for {
		v, ok := lt.contest()
		if !ok {
			break
		}
		rnode.Next = &ListNode{
			Val:  v,
			Next: nil,
		}

		rnode = rnode.Next
	}

	return rhead.Next

}

func adjustNodeList(nodes []*ListNode, size, i int) (newSize int) {
	// find a non-nil node from tail to head
	for newSize = size - 1; newSize > i && nodes[newSize] == nil; newSize-- {
	}

	if newSize <= i {
		return newSize
	}

	nodes[i] = nodes[newSize]
	nodes[newSize] = nil
	return newSize
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 80 ms 5.6 MB
func mergeKLists2(lists []*ListNode) *ListNode {
	rhead := &ListNode{}
	rnode := rhead

	size := len(lists)
	nodes := make([]*ListNode, size)
	copy(nodes, lists)
	for {
		minNodeVal := math.MaxInt64
		minIndex := -1
		for i := 0; i < size; i++ {
			if nodes[i] == nil {
				size = adjustNodeList(nodes, size, i)
				if size <= i {
					break
				}
			}

			if minNodeVal >= nodes[i].Val {
				minNodeVal = nodes[i].Val
				minIndex = i
			}
		}

		if size == 0 || minIndex < 0 {
			break
		}

		rnode.Next = &ListNode{
			Val:  minNodeVal,
			Next: nil,
		}
		rnode = rnode.Next
		nodes[minIndex] = nodes[minIndex].Next
	}

	return rhead.Next
}

func showList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}

	fmt.Println("")
}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	list3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	}

	list4 := &ListNode{
		Val:  0,
		Next: nil,
	}

	list5 := &ListNode{
		Val:  1,
		Next: nil,
	}

	showList(mergeKLists([]*ListNode{list1, list2, list3}))
	showList(mergeKLists([]*ListNode{list5, list4}))
}
