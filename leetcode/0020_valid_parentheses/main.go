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

// https://leetcode.com/problems/valid-parentheses/

var parenthesesMap = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
}

// 0 ms	2 MB
func isValid(s string) bool {
	sLen := len(s)
	if sLen == 0 {
		return true
	}

	top, stack := -1, make([]rune, sLen/2)
	for _, v := range s {
		r, ok := parenthesesMap[v]
		// left pair
		if ok {
			if top < len(stack)-1 {
				top++
				stack[top] = r
			} else {
				return false
			}

			continue
		}
		// right pair
		if top > -1 && stack[top] == v {
			top--
		} else {
			return false
		}
	}

	return top == -1
}

func main() {
	fmt.Printf("r: %t \n", isValid("()[]{}"))
}
