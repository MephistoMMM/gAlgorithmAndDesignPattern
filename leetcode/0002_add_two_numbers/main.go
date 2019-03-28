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

import "fmt"

// https://leetcode.com/problems/add-two-numbers/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	carry := 0

	node := head
	for l1 != nil || l2 != nil {
		val := carry
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		if val >= 10 {
			val %= 10
			carry = 1
		} else {
			carry = 0
		}

		node.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		node = node.Next
	}

	if carry != 0 {
		node.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return head.Next
}

func createListFrom(value []int) *ListNode {
	head := &ListNode{}
	node := head
	for _, v := range value {
		node.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		node = node.Next
	}
	return head.Next
}

func main() {
	// (2 -> 4 -> 3) + (5 -> 6 -> 4)
	v1 := createListFrom([]int{2, 4, 3})
	v2 := createListFrom([]int{5, 6, 4})
	r := addTwoNumbers(v1, v2)
	for r != nil {
		fmt.Printf("%d", r.Val)
		r = r.Next
	}
}
