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

// https://leetcode.com/problems/maximum-subarray/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 4 ms	3.3 MB
func maxSubArray(nums []int) int {
	curSum, maxSum := nums[0], nums[0]
	numsLen := len(nums)

	for i := 1; i < numsLen; i++ {
		curSum = max(nums[i], curSum+nums[i])
		maxSum = max(maxSum, curSum)
	}

	return maxSum
}

// 176 ms	3.5 MB
func maxSubArray2(nums []int) int {
	maxSum := nums[0]
	numsLen := len(nums)

	// sums for subarry long 1
	tmp := make([]int, numsLen)

	for l := 1; l <= numsLen; l++ {
		for i := 0; i < numsLen-l+1; i++ {
			tmp[i] += nums[i+l-1]
			if tmp[i] > maxSum {
				maxSum = tmp[i]
			}
		}
	}

	return maxSum
}

func main() {
	cnsl := &utils.Console{}
	cnsl.Value(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
