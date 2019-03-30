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

// https://leetcode.com/problems/longest-palindromic-substring/

func createAndInitTable(strlen int) [][]bool {
	table := make([][]bool, strlen)
	for i := 0; i < strlen; i++ {
		table[i] = make([]bool, strlen)
		for j := 0; j <= i; j++ {
			table[i][j] = true
		}
	}

	return table
}

// S[i,j] = | true , j > i
//          | false, if S[i+1, j-1] is not palindrome
//          | false, if S[i] != S[j]
//          | j-i+1
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	table := createAndInitTable(len(s))
	max := 1
	substr := [2]int{0, 0}

	for l := 1; l < len(s); l++ {
		for i := 0; i < len(s)-l; i++ {
			j := i + l
			if !table[i+1][j-1] || s[i] != s[j] {
				table[i][j] = false
				continue
			}

			table[i][j] = true
			if j-i+1 > max {
				max = j - i + 1
				substr[0], substr[1] = i, j
			}
		}
	}

	return s[substr[0] : substr[1]+1]
}

func main() {
	s := "babad"
	fmt.Printf("r: %s.", longestPalindrome(s))
}
