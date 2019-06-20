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

// https://leetcode.com/problems/insert-interval/

type OverLapCase uint8

const (
	NONE OverLapCase = 0 + iota
	PARTITION
	CONTAIN
)

func binarySearchPosition(intervals [][]int, newInterval []int) int {
	l, r := 0, len(intervals)
	for l < r {
		m := (l + r) / 2
		if intervals[m][0] < newInterval[0] {
			l = m + 1
		} else if intervals[m][0] > newInterval[0] {
			r = m
		} else {
			// equal
			l = m
			break
		}
	}

	return l
}

func isOverlap(a, b []int) OverLapCase {
	// a: |xxxxxx|
	// b:    |xxxxxx|
	if b[0] <= a[1] && b[1] >= a[1] {
		return PARTITION
	}

	// a: |xxxxxxx|
	// b:  |xxxxx|
	if a[1] >= b[1] {
		return CONTAIN
	}

	return NONE
}

// 4 ms	6.1 MB
func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	startIndex := binarySearchPosition(intervals, newInterval)

	i := 0
	cur := newInterval
	if startIndex != 0 {
		i = startIndex - 1
		cur = intervals[i]
		overlapCase := isOverlap(intervals[i], newInterval)
		if overlapCase == NONE {
			i = startIndex
			cur = newInterval
		}
		if overlapCase == PARTITION {
			cur[1] = newInterval[1]
		}
		if overlapCase == CONTAIN {
			return intervals
		}

		result = append(result, intervals[0:i]...)
		i = startIndex
	}

	for i < len(intervals) {
		overlapCase := isOverlap(cur, intervals[i])
		if overlapCase == PARTITION {
			cur[1] = intervals[i][1]
		}

		if overlapCase == NONE {
			result = append(result, cur)
			cur = nil
			break
		}
		i++
	}

	if cur != nil {
		result = append(result, cur)
	}

	if i != len(intervals) {
		result = append(result, intervals[i:]...)
	}

	return result
}

func main() {
	cnsl := &utils.Console{}
	cnsl.List(insert([][]int{
		{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
	}, []int{6, 8}))
	cnsl.List(insert([][]int{
		{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
	}, []int{1, 8}))
	cnsl.List(insert([][]int{
		{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
	}, []int{13, 17}))
	cnsl.List(insert([][]int{
		{0, 5}, {9, 12},
	}, []int{7, 16}))
}
