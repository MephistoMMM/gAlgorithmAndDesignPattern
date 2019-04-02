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

// https://leetcode.com/problems/longest-common-prefix/

// 0 ms	2.4 MB
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	baseStr := strs[0]
	strs = strs[1:]
	for i := 0; i < len(baseStr); i++ {
		for _, v := range strs {
			if len(v) <= i || v[i] != baseStr[i] {
				return baseStr[:i]
			}
		}
	}

	return baseStr
}

func main() {
	fmt.Printf("r: %s ", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Printf("r: %s ", longestCommonPrefix([]string{"aa", "a"}))

}
