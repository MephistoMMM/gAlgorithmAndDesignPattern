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

import "github.com/MephistoMMM/gAlgorithmAndDesignPattern/leetcode/utils"

// https://leetcode.com/problems/sort-colors/

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

// 0 ms	2.3 MB
func sortColors(nums []int) {
	sLen := len(nums)
	if sLen <= 1 {
		return
	}

	p0, p2 := -1, sLen
	for i := 0; i < p2; {
		if nums[i] == 1 {
			i++
			continue
		}

		if nums[i] == 2 {
			p2--
			swap(nums, i, p2)
			continue
		}

		// nums[i] == 0
		p0++
		if p0 == i {
			i++
		} else {
			swap(nums, i, p0)
		}
	}
}

func main() {
	cnsl := &utils.Console{}
	v := []int{2, 0, 2, 1, 1, 0}
	sortColors(v)
	cnsl.Array(v)
	v1 := []int{1, 2, 0}
	sortColors(v1)
	cnsl.Array(v1)

}
