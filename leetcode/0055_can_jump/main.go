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

// https://leetcode.com/problems/jump-game/

// 8 ms	4.2 MB
func canJump(nums []int) bool {
	numsLen := len(nums)
	for numsLen > 1 {
		i := numsLen - 2
		for i >= 0 && numsLen-i-1 > nums[i] {
			i--
		}

		if i < 0 {
			return false
		}

		numsLen = i + 1
	}

	return true
}

// 8 ms	6.2 MB
func canJump2(nums []int) bool {
	numsLen := len(nums)
	if numsLen <= 1 {
		return true
	}

	i := numsLen - 2
	for i >= 0 && numsLen-i-1 > nums[i] {
		i--
	}

	if i < 0 {
		return false
	}
	return canJump(nums[:i+1])
}

func main() {
	cnls := &utils.Console{}
	cnls.Value(canJump([]int{2, 3, 1, 1, 4}))
	cnls.Value(canJump([]int{3, 2, 1, 0, 4}))
}
