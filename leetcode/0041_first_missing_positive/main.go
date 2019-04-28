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

// https://leetcode.com/problems/first-missing-positive/

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func Negative(v int) int {
	if v > 0 {
		return -v
	}
	return v
}

func packupPositive(nums []int) int {
	p := -1
	for i, numsLen := 0, len(nums); i < numsLen; i++ {
		if nums[i] > 0 {
			p++
			tmp := nums[p]
			nums[p] = nums[i]
			nums[i] = tmp
		}
	}

	return p
}

// 0 ms	2.2 MB
func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	k := packupPositive(nums) + 1
	for i := 0; i < k; i++ {
		v := Abs(nums[i])
		if v <= k {
			nums[v-1] = Negative(nums[v-1])
		}
	}

	firstPositiveIndex := k
	for i := 0; i < k; i++ {
		if nums[i] > 0 {
			firstPositiveIndex = i
			break
		}
	}

	return firstPositiveIndex + 1
}

func main() {
	fmt.Printf("r: %d \n", firstMissingPositive([]int{1, 2, 0}))
	fmt.Printf("r: %d \n", firstMissingPositive([]int{3, 4, -1, 1}))
}
