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

import "fmt"

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

// 8 ms	4.1 MB
func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}

	// first binary search find leftest
	l0, r0 := 0, len(nums)
	for l0 != r0 {
		m := (l0 + r0) / 2
		if nums[m] >= target {
			r0 = m
		} else {
			l0 = m + 1
		}
	}
	if l0 == len(nums) || nums[l0] != target {
		return result
	}
	result[0] = l0

	// second binary search find rightest
	l1, r1 := l0, len(nums)
	for l1 != r1 {
		m := (l1 + r1) / 2
		if nums[m] <= target {
			l1 = m + 1
		} else {
			r1 = m
		}
	}
	result[1] = r1 - 1

	return result
}

func binarySearch(nums []int, l, r *int, target int) (m int) {
	for *l <= *r {
		m = (*l + *r) / 2
		if nums[m] == target {
			break
		}

		if nums[m] > target {
			*r = m - 1
		} else {
			*l = m + 1
		}
	}

	if *l > *r {
		return -1
	}

	return m
}

//8 ms	4.1 MB
func searchRange2(nums []int, target int) []int {
	result := []int{-1, -1}
	// first binary search find a target
	l0, r0 := 0, len(nums)-1
	m0 := binarySearch(nums, &l0, &r0, target)
	if m0 < 0 {
		return result
	}

	result[0], result[1] = m0, m0
	// second find min index
	l, r := l0, m0-1
	m1 := 0
	for {
		m1 = binarySearch(nums, &l, &r, target)
		if m1 < 0 {
			break
		} else {
			result[0] = m1
			r = m1 - 1
		}
	}
	// third find max index
	l, r = m0+1, r0
	for {
		m1 = binarySearch(nums, &l, &r, target)
		if m1 < 0 {
			break
		} else {
			result[1] = m1
			l = m1 + 1
		}
	}

	return result
}

func main() {
	fmt.Printf("%v \n", searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Printf("%v \n", searchRange([]int{}, 8))
	fmt.Printf("%v \n", searchRange([]int{1}, 8))
	fmt.Printf("%v \n", searchRange([]int{8}, 8))
	fmt.Printf("%v \n", searchRange([]int{1, 8}, 8))
	fmt.Printf("%v \n", searchRange([]int{8, 8}, 8))
	fmt.Printf("%v \n", searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}
