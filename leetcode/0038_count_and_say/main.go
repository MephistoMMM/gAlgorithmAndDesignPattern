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

// https://leetcode.com/problems/count-and-say/

import (
	"fmt"
	"strconv"
)

func countAndRead(c string) string {
	result := make([]byte, 0, len(c))
	for len(c) > 0 {
		count := 0
		prev := c[0]
		for len(c) > count+1 && c[count+1] == prev {
			count++
		}

		result = append(result, []byte(strconv.Itoa(count+1))...)
		result = append(result, byte(prev))
		c = c[count+1:]
	}

	return string(result)
}

func countAndSay(n int) string {
	cur := "1"
	for i := 1; i < n; i++ {
		cur = countAndRead(cur)
	}

	return cur
}

func main() {
	fmt.Printf("r: %s\n", countAndSay(5))
}
