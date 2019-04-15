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

// https://leetcode.com/problems/search-insert-position/

//4 ms 3.1 MB
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	for l != r {
		m := (l + r) / 2
		if nums[m] == target {
			r = m
			l = m
			break
		}

		if nums[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}

	return r
}

func main() {
	fmt.Printf("r: %d \n", searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Printf("r: %d \n", searchInsert([]int{1, 3, 5, 6}, 0))
	fmt.Printf("r: %d \n", searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Printf("r: %d \n", searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Printf("r: %d \n", searchInsert([]int{}, 2))
	fmt.Printf("r: %d \n", searchInsert([]int{1}, 2))
	fmt.Printf("r: %d \n", searchInsert([]int{2}, 2))
}
