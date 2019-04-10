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

// https://leetcode.com/problems/next-permutation/solution/

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func reverse(nums []int) {
	nLen := len(nums)
	for i := 0; i < nLen/2; i++ {
		swap(nums, i, nLen-1-i)
	}
}

// 8 ms 3.4 MB
func nextPermutation(nums []int) {
	nLen := len(nums)
	if nLen < 2 {
		return
	}

	i := nLen - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := nLen - 1
		for j >= i && nums[j] <= nums[i] {
			j--
		}
		swap(nums, i, j)
	}
	reverse(nums[i+1:])
}

func main() {
	nums := []int{1, 5, 8, 4, 7, 6, 5, 3, 1}
	nextPermutation(nums)
	fmt.Printf("r: %v \n", nums)
	nums2 := []int{3, 2, 1}
	nextPermutation(nums2)
	fmt.Printf("r: %v \n", nums2)
}
