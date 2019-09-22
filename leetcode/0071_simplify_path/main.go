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
	"strings"
)

// https://leetcode.com/problems/simplify-path/

// 0 ms	2.8 MB
func simplifyPath(path string) string {
	if len(path) == 0 {
		return path
	}

	stack := make([]string, 0)
	top := -1

	prev, i := 0, 1
	for ; i < len(path); i++ {
		if path[i] != '/' {
			if i == len(path)-1 {
				i++
			} else {
				continue
			}
		}

		directory := path[prev+1 : i]
		if len(directory) == 0 || directory == "." {
			// do nothing
		} else if directory == ".." && top < 0 {
			// do nothing
		} else if directory == ".." && top >= 0 {
			stack = stack[:top]
			top--
		} else {
			stack = append(stack, path[prev+1:i])
			top++
			prev = i
		}

		prev = i
	}

	return "/" + strings.Join(stack, "/")

}

func main() {
	cnsl := &utils.Console{}
	cnsl.Value(simplifyPath("/home/"))
	cnsl.Value(simplifyPath("/a//b////c/d//././/.."))
	cnsl.Value(simplifyPath("/../"))
}
