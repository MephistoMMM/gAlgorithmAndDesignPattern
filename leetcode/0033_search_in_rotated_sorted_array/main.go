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
)

// https://leetcode.com/problems/search-in-rotated-sorted-array/

// binary search approach: https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/14425/Concise-O(log-N)-Binary-search-solution

// 4 ms 2.6 MB
func search(nums []int, target int) int {
	nLen := len(nums)
	if nLen == 0 {
		return -1
	}

	i := 0
	for i >= 0 {
		if nums[i] == target {
			return i
		}

		next := 0
		if nums[i] < target {
			next = (i + 1 + nLen) % nLen
			// go to right end
			if nums[i] >= nums[next] {
				break
			}
		}

		if nums[i] > target {
			next = (i - 1 + nLen) % nLen
			// go to left end
			if nums[i] <= nums[next] {
				break
			}
		}

		dis := nums[i] - target
		nextDis := nums[next] - target
		// find failed
		if dis*nextDis < 0 {
			break
		}

		i = next
	}

	if nums[i] != target {
		return -1
	}

	return i
}

func main() {
	fmt.Printf("r: %d \n", search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Printf("r: %d \n", search([]int{4, 5, 6, 7, 0, 1, 2}, 8))
	fmt.Printf("r: %d \n", search([]int{4}, 4))
}
