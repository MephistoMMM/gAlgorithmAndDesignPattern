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
	"github.com/MephistoMMM/gAlgorithmAndDesignPattern/leetcode/utils"
)

// https://leetcode.com/problems/minimum-window-substring/

// AlphaDict 0~25 : A ~ Z
type AlphaDict struct {
	length int
	dict   [128]int
}

func (ad *AlphaDict) Len() int {
	return ad.length
}

func (ad *AlphaDict) Contain(c byte) bool {
	return ad.dict[c] > 0
}

func (ad *AlphaDict) Add(c byte) {
	ad.length++
	ad.dict[c]++
}

func (ad *AlphaDict) Sub(c byte) {
	if ad.Contain(c) {
		ad.length--
		ad.dict[c]--
	}
}

// AlphaLinkListNode :
type AlphaLinkListNode struct {
	// main link prev and next point
	Prev *AlphaLinkListNode
	Next *AlphaLinkListNode

	// same character link next point
	SCNext *AlphaLinkListNode

	Char  byte
	Index int
}

type SameCharLinkList struct {
	head *AlphaLinkListNode
	tail *AlphaLinkListNode
}

func (scll *SameCharLinkList) Shift() {
	if scll.head == nil {
		return
	}
	if scll.head == scll.tail {
		scll.head = nil
		scll.tail = nil
		return
	}
	scll.head = scll.head.SCNext
}

func (scll *SameCharLinkList) Append(node *AlphaLinkListNode) {
	if scll.head == nil {
		scll.head = node
		scll.tail = node
		return
	}

	scll.tail.SCNext = node
	scll.tail = node
}

// AlphaLinkList :
type AlphaLinkList struct {
	head *AlphaLinkListNode
	tail *AlphaLinkListNode

	charsList [128]SameCharLinkList
}

func (all *AlphaLinkList) HeadIndex() int {
	// head == NULL
	if all.head == nil {
		return -1
	}
	// head != null
	return all.head.Index
}

func (all *AlphaLinkList) TailIndex() int {
	// tail == NULL
	if all.tail == nil {
		return -1
	}
	return all.tail.Index
}

func (all *AlphaLinkList) Len() int {
	// head == NULL
	if all.head == nil {
		return -1
	}
	// head != null
	return all.tail.Index - all.head.Index + 1
}

func (all *AlphaLinkList) Remove(c byte) {
	node := all.charsList[int(c)].head
	if node == nil {
		return
	}

	if node == all.head {
		// head = cur->next
		all.head = node.Next
	} else {
		// prev->next = cur->next
		node.Prev.Next = node.Next
	}

	if node == all.tail {
		// tail = cur->prev
		all.tail = node.Prev
	} else {
		// next->prev = cur->prev
		node.Next.Prev = node.Prev
	}

	all.charsList[int(c)].Shift()
}

func (all *AlphaLinkList) Append(c byte, index int) {
	i := int(c)
	node := &AlphaLinkListNode{
		Char:  c,
		Index: index,
	}
	all.charsList[i].Append(node)

	// head is NULL
	if all.head == nil {
		// head = cur
		all.head = node
		// tail = cur
		all.tail = node
		return
	}

	// tail->next = cur
	all.tail.Next = node
	// cur->prev = tail
	node.Prev = all.tail
	// tail = cur
	all.tail = node
	return
}

// 16 ms 6.4 MB
func minWindow(s string, t string) string {
	if t == "" || s == "" {
		return ""
	}

	summary := &AlphaDict{}
	for _, c := range []byte(t) {
		summary.Add(c)
	}

	minLen := len(s) + 1
	minStr := ""
	linkList := &AlphaLinkList{}
	status := *summary
	for index, c := range []byte(s) {
		if !summary.Contain(c) {
			continue
		}

		if !status.Contain(c) {
			linkList.Remove(c)
		}

		status.Sub(c)
		linkList.Append(c, index)
		if status.Len() == 0 && linkList.Len() < minLen {
			minLen = linkList.Len()
			minStr = s[linkList.HeadIndex() : linkList.TailIndex()+1]
		}

	}

	return minStr
}

func main() {
	cnsl := &utils.Console{}
	cnsl.Value(minWindow("acbbaca", "aba"))
	cnsl.Value(minWindow("ADOBECODEBANC", "ABC"))
	cnsl.Value(minWindow("a", "a"))
	cnsl.Value(minWindow("a", "aa"))
}
