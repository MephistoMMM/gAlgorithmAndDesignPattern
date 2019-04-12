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

// https://leetcode.com/problems/longest-valid-parentheses/

func createTable(sLen int) [][]bool {
	table := make([][]bool, sLen)
	for i := range table {
		table[i] = make([]bool, sLen/2+1)
		table[i][0] = true
	}

	return table
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 4 ms	2.4 MB
func longestValidParentheses(s string) int {
	sLen := len(s)
	if sLen < 2 {
		return 0
	}

	left, right, max := 0, 0, 0
	for i := 0; i < sLen; i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			max = maxInt(max, right*2)
		}

		if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := sLen - 1; i > -1; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			max = maxInt(max, left*2)
		}

		if right < left {
			left, right = 0, 0
		}
	}

	return max
}

// 0 ms 3.2 MB
func longestValidParenthesesOn(s string) int {
	sLen := len(s)
	if sLen < 2 {
		return 0
	}

	record := make([]int, sLen+1)
	stack := make([]int, sLen)
	top := -1
	for i := range s {
		if s[i] == ')' {
			if top > -1 {
				record[stack[top]] = i
				top--
			}
		} else {
			top++
			stack[top] = i
		}
	}

	max := 0
	count := 0
	for i := 0; i <= sLen; i++ {
		if record[i] == 0 {
			if count > max {
				max = count
			}
			count = 0
			continue
		}

		count += record[i] - i + 1
		i = record[i]
	}

	return max
}

// t[i,l] = | true , if s[i] == '(' and s[i+l-1] == ')' and t[i+1, l-1]
//          | true , if t[i, k] && t[i+k, l-k]
//          | else false
func longestValidParenthesesSlow(s string) int {
	sLen := len(s)
	if sLen < 2 {
		return 0
	}

	table := createTable(sLen)

	max := 0
	for l := 1; l <= sLen/2; l++ {
		for i := 0; i < sLen-l*2+1; i++ {
			cur := table[i+1][l-1]
			if cur && !(s[i] == '(' && s[i+l*2-1] == ')') {
				cur = false
			}

			if !cur {
				for k := 1; k < l; k++ {
					cur = table[i][k] && table[i+k*2][l-k]
					if cur {
						break
					}
				}
			}

			table[i][l] = cur
			if cur {
				max = l * 2
			}
		}
	}

	return max
}

func main() {
	fmt.Printf("r: %d \n", longestValidParentheses("(()"))
	fmt.Printf("r: %d \n", longestValidParentheses(")()())"))
	fmt.Printf("r: %d \n", longestValidParentheses(")(((((()())()()))()(()))("))
	fmt.Printf("r: %d \n", longestValidParentheses("()(()"))
}
