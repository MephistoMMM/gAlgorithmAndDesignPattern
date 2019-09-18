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

import "gAaD/leetcode/utils"

// https://leetcode.com/problems/minimum-window-substring/

// AlphaDict 0~25 : A ~ Z
type AlphaDict struct {
	length int
	dict   [128]bool
}

func (ad *AlphaDict) Len() int {
	return ad.length
}

func (ad *AlphaDict) Contain(c byte) bool {
	return ad.dict[int(c)]
}

func (ad *AlphaDict) Add(c byte) {
	if !ad.Contain(c) {
		ad.length++
		ad.dict[int(c)] = true
	}
}

func (ad *AlphaDict) Remove(c byte) {
	if ad.Contain(c) {
		ad.length--
		ad.dict[int(c)] = false
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

// AlphaLinkList :
type AlphaLinkList struct {
	head *AlphaLinkListNode
	tail *AlphaLinkListNode

	charsList []*AlphaLinkList
}

func NewAlphaLinkList() *AlphaLinkList {
	charsList := make([]*AlphaLinkList, 128)
	for i := range charsList {
		charsList[i] = &AlphaLinkList{}
	}

	return &AlphaLinkList{
		charsList: charsList,
	}

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

func (all *AlphaLinkList) IsHead(c byte) bool {
	return all.head.Char == c
}

func (all *AlphaLinkList) IsTail(c byte) bool {
	return all.tail.Char == c
}

func (all *AlphaLinkList) Delete(c byte) {
	if all.IsHead(c) {
		// header = cur->next
		all[1][0] = all[i][1]
	} else {
		// prev->next = cur->next
		all[all[i][0]][1] = all[i][1]
	}

	if all.IsTail(c) {
		// tailer = cur->prev
		all[1][1] = all[i][0]
	} else {
		// next->prev = cur->prev
		all[all[i][1]][0] = all[i][0]
	}
	all[i][1] = 0
	all[i][0] = 0
}

func (all *AlphaLinkList) Append(c byte, index int) {
	i := int(c)
	// header is NULL
	if all[1][0] == 0 {
		// header = cur
		all[1][0] = i
		// tailer = cur
		all[1][1] = i
		return
	}

	// tailer->next = cur
	all[all[1][1]][1] = i
	// cur->prev = tailer
	all[i][0] = all[1][1]
	// tai
	all[1][1] = i
	// cur->value = index
	all[i][2] = index
	return
}

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

		status.Remove(c)
		linkList.Append(c, index)
		if status.Len() == 0 && linkList.Len() < minLen {
			minLen = linkList.Len()
			minStr = s[linkList.HeaderIndex() : linkList.TailerValue()+1]
		}

	}

	return minStr
}

func main() {
	cnsl := &utils.Console{}
	cnsl.Value(minWindow("ADOBECODEBANC", "ABC"))
	cnsl.Value(minWindow("a", "a"))

}
