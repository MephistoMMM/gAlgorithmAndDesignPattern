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
)

// https://leetcode.com/problems/edit-distance/

func createRecordTable(w1Len, w2Len int) [][]int {
	recordTable := make([][]int, w1Len+1)
	for i := range recordTable {
		recordTable[i] = make([]int, w2Len+1)
	}
	// if w2 is empty, action(delete) always equal to length of w1
	for i := 0; i <= w1Len; i++ {
		recordTable[i][0] = i
	}
	// if w1 is empty, action(add) always equal to length of w2
	for j := 0; j <= w2Len; j++ {
		recordTable[0][j] = j
	}

	return recordTable
}

func min(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}

	return m
}

// 4 ms	2.6 MB
//
// Wonderful: https://leetcode.com/problems/edit-distance/discuss/25911/My-O(mn)-time-and-O(n)-space-solution-using-DP-with-explanation
func minDistanceN(word1 string, word2 string) int {
	w1Len, w2Len := len(word1), len(word2)
	if w1Len == 0 {
		return w2Len
	}
	if w2Len == 0 {
		return w1Len
	}

	rt := make([]int, w2Len+1)
	for j := 0; j < w2Len+1; j++ {
		rt[j] = j
	}

	for i := 1; i <= w1Len; i++ {
		prev := i
		for j := 1; j <= w2Len; j++ {
			cur := rt[j-1]
			if word1[i-1] != word2[j-1] {
				cur = 1 + min(rt[j], prev, rt[j-1])
			}

			rt[j-1] = prev
			prev = cur

		}
		rt[w2Len] = prev
	}

	return rt[w2Len]
}

// 4 ms	5.6 MB
func minDistance(word1 string, word2 string) int {
	w1Len, w2Len := len(word1), len(word2)
	if w1Len == 0 {
		return w2Len
	}
	if w2Len == 0 {
		return w1Len
	}

	rt := createRecordTable(w1Len, w2Len)

	for i := 1; i <= w1Len; i++ {
		for j := 1; j <= w2Len; j++ {
			if word1[i-1] == word2[j-1] {
				rt[i][j] = rt[i-1][j-1]
				continue
			}

			rt[i][j] = 1 + min(rt[i-1][j-1], rt[i][j-1], rt[i-1][j])
		}
	}

	return rt[w1Len][w2Len]
}

func showRecordTable(rt [][]int) {
	fmt.Println("- - -")
	for _, v := range rt {
		fmt.Printf("[ ")
		for _, k := range v {
			if k < 0 {
				fmt.Printf("_ ")
				continue
			}

			fmt.Printf("%d ", k)
		}
		fmt.Printf("]\n")
	}
}

func main() {
	fmt.Printf("r: %d \n", minDistance("intention", "execution"))
	fmt.Printf("r: %d \n", minDistance("horse", "ros"))
	fmt.Printf("r: %d \n", minDistance("a", "b"))
	fmt.Printf("r: %d \n", minDistanceN("kitten", "sitting"))
	fmt.Printf("r: %d \n", minDistanceN("intention", "execution"))
}
