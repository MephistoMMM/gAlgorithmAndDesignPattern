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

// https://leetcode.com/problems/permutation-sequence/

func factorial(n int) int {
	result := 1
	for i := n; i > 1; i-- {
		result *= i
	}
	return result
}

// 0 ms	2 MB
func getPermutation(n int, k int) string {
	k -= 1
	usedNums := make([]bool, n)
	result := make([]byte, 0, n)
	for i := n - 1; i >= 0; i-- {
		divisor := factorial(i)
		quotient := k / divisor
		j := 0
		for ; j < n; j++ {
			if !usedNums[j] {
				quotient--
			}
			if quotient < 0 {
				break
			}
		}
		usedNums[j] = true
		result = append(result, '1'+byte(j))

		k = k % divisor
	}

	return string(result)
}

func main() {
	cnsl := &utils.Console{}
	cnsl.Value(getPermutation(4, 9))
}
