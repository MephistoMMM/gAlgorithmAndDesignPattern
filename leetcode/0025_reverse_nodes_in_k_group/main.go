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

// https://leetcode.com/problems/reverse-nodes-in-k-group/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseK(head *ListNode, k int) (result, residual *ListNode) {
	var rnode *ListNode
	i := 0
	for ; i < k && head != nil; i++ {
		node := head
		head = head.Next
		node.Next = rnode
		rnode = node
	}

	if i == k {
		return rnode, head
	}

	var rrnode *ListNode
	for rnode != nil {
		node := rnode
		rnode = rnode.Next
		node.Next = rrnode
		rrnode = node
	}

	return rrnode, nil
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	vhead := &ListNode{Next: head}
	kprev := vhead

	for kprev.Next != nil {
		rnodes, residual := reverseK(kprev.Next, k)
		if residual == nil {
			kprev.Next = rnodes
			break
		}

		tmp := kprev.Next
		kprev.Next = rnodes
		tmp.Next = residual
		kprev = tmp
	}

	return vhead.Next
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	vhead := &ListNode{Next: head}
	cur, kprev := vhead, vhead

	for cur != nil {
		for i := k; i > 0 && cur != nil; i-- {
			cur = cur.Next
		}

		if cur == nil {
			break
		}

		tnode := cur.Next
		p := kprev.Next
		for {
			node := p
			p = p.Next
			node.Next = tnode
			tnode = node

			if node == cur {
				break
			}
		}

		tmp := kprev.Next
		kprev.Next = tnode
		kprev = tmp
		cur = kprev
	}

	return vhead.Next
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
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	// showList(reverseKGroup(list1, 2))
	showList(reverseKGroup(list1, 3))
}
