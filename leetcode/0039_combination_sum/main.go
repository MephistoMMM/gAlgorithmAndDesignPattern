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

// https://leetcode.com/problems/combination-sum/

// 4 ms	4 MB
func backTrackCombinationSum(candidates []int, current *[]int, target int, result *[][]int) {
	if len(candidates) == 0 {
		return
	}

	if target == 0 {
		r := make([]int, len(*current))
		copy(r, *current)
		*result = append(*result, r)
		return
	}

	curLen := len(*current)
	for i := range candidates {
		if candidates[i] > target {
			break
		}
		*current = append(*current, candidates[i])
		backTrackCombinationSum(candidates[i:], current, target-candidates[i], result)
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

// 8 ms	3.9 MB
// return boolean to decide whether continue
func backTrackCombinationSum2(candidates []int, current *[]int, target int, result *[][]int) {
	if len(candidates) == 0 {
		return
	}

	if candidates[0] > target {
		return
	}

	curLen := len(*current)
	if curLen == 0 || (*current)[curLen-1] != candidates[0] {
		// ignore value
		backTrackCombinationSum(candidates[1:], current, target, result)
		*current = (*current)[:curLen]
	}

	*current = append(*current, candidates[0])
	target -= candidates[0]
	curLen = len(*current)
	if target == 0 {
		r := make([]int, curLen)
		copy(r, *current)
		*result = append(*result, r)
	}

	// current value
	backTrackCombinationSum(candidates, current, target, result)
	*current = (*current)[:curLen]
	// next value
	backTrackCombinationSum(candidates[1:], current, target, result)
	*current = (*current)[:curLen]
}

func combinationSum2(candidates []int, target int) [][]int {
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
	fmt.Printf("}\n")
}

func main() {
	showResult(combinationSum([]int{2, 3, 6, 7}, 7))
	showResult(combinationSum([]int{2, 3, 5}, 7))
}
