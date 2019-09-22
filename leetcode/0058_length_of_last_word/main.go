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

import "github.com/MephistoMMM/gAlgorithmAndDesignPattern/leetcode/utils"

// https://leetcode.com/problems/length-of-last-word/

// 0 ms	2.2 MB
func lengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i > -1; i-- {
		if s[i] != ' ' {
			count++
		} else if count != 0 {
			break
		}
		// do nothing if s[i] == ' ' and count == 0
	}

	return count
}

func main() {
	cnsl := &utils.Console{}
	cnsl.Value(lengthOfLastWord("Hello World"))
}
