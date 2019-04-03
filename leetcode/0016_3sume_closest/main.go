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
	"math"
	"sort"
)

// https://leetcode.com/problems/3sum-closest/

func distance(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// 4 ms	2.7 MB
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	min := nums[0] + nums[1] + nums[2]
	di := math.MaxInt64
	for i := 0; i < len(nums)-2; i++ {
		diI := nums[i] - target
		for l, r := i+1, len(nums)-1; l < r; {
			// omit wrong solution
			if (diI <= 0 && nums[r] < 0) || (diI >= 0 && nums[l] > 0) {
				break
			}

			sum := nums[l] + nums[r] + nums[i]
			if sum == target {
				return sum
			}

			diSum := distance(sum, target)
			if diSum <= di {
				min = sum
				di = diSum
			}

			// omit duplicate number
			if sum < target {
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			} else {
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}

		}
	}

	return min
}

func main() {
	fmt.Printf("r: %d \n", threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Printf("r: %d \n", threeSumClosest([]int{1, 1, -1, -1, 3}, 3))
	fmt.Printf("r: %d \n", threeSumClosest([]int{1, 1, 1, 1, 0}, 100))
	fmt.Printf("r: %d \n", threeSumClosest([]int{0, 0, 0}, 1))
	fmt.Printf("r: %d \n", threeSumClosest([]int{0, 1, 2}, 0))
}
