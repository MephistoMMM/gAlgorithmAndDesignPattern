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

// https://leetcode.com/problems/3sum/

// 892 ms 272.9 MB
// 大概思路是外围For一遍数组，然后在每次迭代的时候，设置[l,r]的区间，区间范围为l , r = i + 1, len(nums) - 1, 这样比对的数就有3个，分别是nums[i], nums[l] 和 nums[r].
// 最终要达到的效果是: nums[l] + nums[r] == -nums[i]
// 这道题一定要记住去重，不仅仅是区间的l和r要去重，外围的i也需要去重。去重的方法如下:
// i去重： if i > 0 or nums[i] == nums[i-1]
// l去重： while l < r and nums[l] == nums[l-1]: l += 1
// r去重： while l < r and nums[r] == nums[r+1]: r -= 1
// i中断： nums[i] > 0
// i,l中断： nums[i] + nums[l] > 0
func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}

	// first sort it
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i] > 0 {
			break
		}

		for l, r := i+1, len(nums)-1; l < r; {
			if nums[i]+nums[l] > 0 {
				break
			}

			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				result = append(result,
					[]int{nums[i], nums[l], nums[r]})
			}

			if sum >= 0 {
				r--
				for r == len(nums)-1 || (r > l && nums[r] == nums[r+1]) {
					r--
				}
			}

			if sum <= 0 {
				l++
				for l == i+1 || (l < r && nums[l] == nums[l-1]) {
					l++
				}
			}
		}
	}

	return result
}

func showTuple(tuples [][]int) {
	for _, tuple := range tuples {
		fmt.Printf("%v\t", tuple)
	}
	fmt.Println("")
}

func main() {
	showTuple(threeSum([]int{-1, 0, 1, 2, -1, -4}))

}
