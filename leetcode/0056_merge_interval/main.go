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
	"gAaD/leetcode/utils"
	"sort"
)

// https://leetcode.com/problems/merge-intervals/

// 8 ms	4.9 MB
func merge(intervals [][]int) [][]int {
	result := [][]int{}
	if len(intervals) == 0 {
		return result
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		isOverlap := false
		// |xxxxxx|
		//     |xxxxxx|
		if intervals[i][0] <= cur[1] && intervals[i][1] >= cur[1] {
			cur[1] = intervals[i][1]
			isOverlap = true
		}

		// |xxxxxxx|
		//  |xxxxx|
		if cur[1] >= intervals[i][1] {
			isOverlap = true
		}

		if !isOverlap {
			result = append(result, cur)
			cur = intervals[i]
		}
	}
	result = append(result, cur)

	return result
}

func main() {
	cnsl := &utils.Console{}
	cnsl.List(merge([][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}))
	cnsl.List(merge([][]int{
		{1, 4}, {4, 5},
	}))
	cnsl.List(merge([][]int{
		{1, 4}, {0, 5},
	}))
	cnsl.List(merge([][]int{
		{1, 4}, {2, 3},
	}))
	cnsl.List(merge([][]int{
		{1, 4}, {2, 3}, {7, 9}, {0, 5}, {9, 12},
	}))
}
