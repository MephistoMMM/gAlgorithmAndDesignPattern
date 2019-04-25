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
	"sort"
)

// https://leetcode.com/problems/combination-sum-ii/

// 0 ms	2.6 MB
func backTrackCombinationSum(candidates []int, current *[]int, target int, result *[][]int) {

	if target == 0 {
		r := make([]int, len(*current))
		copy(r, *current)
		*result = append(*result, r)
		return
	}

	candidatesLen := len(candidates)
	if candidatesLen == 0 {
		return
	}

	curLen := len(*current)
	prev := -1
	for i := 0; i < candidatesLen; i++ {
		if candidates[i] > target {
			break
		}
		if prev == candidates[i] {
			continue
		} else {
			prev = candidates[i]
		}

		*current = append(*current, candidates[i])
		backTrackCombinationSum(candidates[i+1:], current, target-candidates[i], result)
		*current = (*current)[:curLen]
	}

}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)
	current := make([]int, 0)
	backTrackCombinationSum(candidates, &current, target, &result)
	return result
}

func showResult(r [][]int) {
	fmt.Printf("{\n")
	for _, v := range r {
		fmt.Printf("\t%v\n", v)
	}

}

func main() {
	showResult(combinationSum([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	showResult(combinationSum([]int{2, 5, 2, 1, 2}, 5))
}
