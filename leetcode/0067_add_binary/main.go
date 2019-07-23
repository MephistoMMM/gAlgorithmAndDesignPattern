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

// https://leetcode.com/problems/add-binary/

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 0 ms	2.2 MB
func addBinary(a string, b string) string {
	maxLen := maxInt(len(a), len(b))
	c := make([]byte, maxLen+1)
	carry := byte(0)

	i, j := len(a)-1, len(b)-1
	for ; i > -1 && j > -1; i, j = i-1, j-1 {
		v := a[i] + b[j] - '0' - '0' + carry
		if v > byte(1) {
			v -= 2
			carry = 1
		} else {
			carry = 0
		}
		c[maxLen] = v + '0'
		maxLen--
	}
	for i > -1 {
		v := a[i] + carry - '0'
		if v > byte(1) {
			v -= 2
			carry = 1
		} else {
			carry = 0
		}
		c[maxLen] = v + '0'
		maxLen--
		i--
	}
	for j > -1 {
		v := b[j] + carry - '0'
		if v > byte(1) {
			v -= 2
			carry = 1
		} else {
			carry = 0
		}
		c[maxLen] = v + '0'
		maxLen--
		j--
	}

	if carry == 1 {
		c[0] = '1'
	} else {
		c = c[1:]
	}

	return string(c)
}

func main() {
	cnsl := &utils.Console{}
	cnsl.Value(addBinary("11", "1"))
}
