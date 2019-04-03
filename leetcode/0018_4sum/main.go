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

// https://leetcode.com/problems/4sum/

func twoSumForSortedArr(nums []int, target int) [][]int {
	result := [][]int{}
	if len(nums) < 2 {
		return result
	}

	for l, r := 0, len(nums)-1; l < r; {
		if nums[l] > 0 && nums[l] > target {
			break
		}

		sum := nums[l] + nums[r]
		if sum == target {
			result = append(result,
				[]int{nums[l], nums[r]})
		}

		if sum >= target {
			r--
			for r > l && nums[r] == nums[r+1] {
				r--
			}
		}

		if sum <= target {
			l++
			for l < r && nums[l] == nums[l-1] {
				l++
			}
		}
	}

	return result
}

func NSumForSortedArr(n int, nums []int, target int) [][]int {
	if n < 2 {
		panic("Parameter N must be greater than 2")
	}

	if n == 2 {
		return twoSumForSortedArr(nums, target)
	}

	result := [][]int{}
	if len(nums) < n {
		return result
	}

	for i := 0; i < len(nums)-n+1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i] > 0 && nums[i] > target {
			break
		}

		for _, v := range NSumForSortedArr(n-1, nums[i+1:], target-nums[i]) {
			result = append(result, append([]int{nums[i]}, v...))
		}
	}

	return result
}

// 12 ms	3.3 MB
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return NSumForSortedArr(4, nums, target)
}

func showTuple(tuples [][]int) {
	for _, tuple := range tuples {
		fmt.Printf("%v\t", tuple)
	}
	fmt.Println("")
}

func main() {
	showTuple(fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11))
	showTuple(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
