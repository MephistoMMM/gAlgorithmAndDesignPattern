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

import "github.com/MephistoMMM/gAlgorithmAndDesignPattern/leetcode/utils"

// https://leetcode.com/problems/rotate-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

// Next ...
func (ln *ListNode) NextNode() utils.LinkListNode {
	return ln.Next
}

// Value ...
func (ln *ListNode) Value() interface{} {
	return ln.Val
}

// 0 ms	2.5 MB
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	// get the length of list and make it be cycle list
	length := 1
	prev, cur := head, head.Next
	for cur != nil {
		length++
		prev, cur = cur, cur.Next
	}
	prev.Next = head

	// convert right rotating to left rotating
	k = length - k%length
	for k > 0 {
		prev = prev.Next
		k--
	}

	head = prev.Next
	prev.Next = nil
	return head
}

func main() {
	cnsl := &utils.Console{}
	head := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4, &ListNode{5, nil},
				},
			},
		},
	}
	cnsl.LinkList(rotateRight(head, 2))
}
