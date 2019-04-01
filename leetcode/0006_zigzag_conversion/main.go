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

// https://leetcode.com/problems/zigzag-conversion/submissions/

import "fmt"

// row\i
//       P     I    N
//       A   L S  I G
//       Y A   H R
//       P     I
//
// row and i start from 0
// n is numRows
func zigzagIndexMap(n int, row int, i int) int {
	if n == 1 {
		return i
	}

	if row == 0 || row == n-1 {
		return row + i*2*(n-1)
	}

	// row + ( (n-row-1) + (n-2-row+1) ) * ((i+1)/2) + (2*i)*(j/2)
	return row + ((i+1)/2)*2*(n-row-1) + (i/2)*2*row

}

func convert(s string, numRows int) string {
	byts := make([]byte, len(s))

	for row := 0; row < numRows; row++ {
		for i := 0; ; i++ {
			k := zigzagIndexMap(numRows, row, i)
			if k >= len(s) {
				break
			}

			byts = append(byts, s[k])
		}
	}

	return string(byts)
}

func main() {
	fmt.Println(convert("AB", 1))
}
