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

// https://leetcode.com/problems/combinations/

// 144ms 218MB
func combine(n int, k int) [][]int {
	if k == 0 || n == 0 {
		return [][]int{}
	}

	results := make([][]int, 0, utils.Combination(k, n))
	cache := make([]int, k)

	cur := 0
	cache[cur] = 1
	for {
		if cache[cur] > n-k+cur+1 {
			if cur == 0 {
				break
			}
			cur--
			cache[cur]++
			continue
		}

		if cur == k-1 {
			v := make([]int, k)
			copy(v, cache)
			results = append(results, v)
			cache[cur]++
			continue
		}

		cur++
		cache[cur] = cache[cur-1] + 1
	}

	return results
}

func main() {
	cnsl := &utils.Console{}
	cnsl.DoubleDimArray(combine(4, 2))
	cnsl.DoubleDimArray(combine(4, 1))
	cnsl.DoubleDimArray(combine(0, 1))
	cnsl.DoubleDimArray(combine(4, 0))
}
