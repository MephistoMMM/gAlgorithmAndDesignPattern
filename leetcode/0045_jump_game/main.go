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

// https://leetcode.com/problems/jump-game-ii/

type UpperTriangularMatrix struct {
	row  int
	elem []int
}

func NewUpperTriangularMatrix(row int) *UpperTriangularMatrix {
	utm := &UpperTriangularMatrix{
		row:  row,
		elem: make([]int, (1+row)*row/2),
	}
	return utm
}

func (utm *UpperTriangularMatrix) index(i, j int) int {
	if i >= utm.row || j < i || j >= utm.row {
		panic("index out of range")
	}

	skip := (utm.row + utm.row - i + 1) * i / 2
	return skip + j - i
}

func (utm *UpperTriangularMatrix) Get(i, j int) int {
	return utm.elem[utm.index(i, j)]
}

// Set ...
func (utm *UpperTriangularMatrix) Set(i, j int, v int) {
	utm.elem[utm.index(i, j)] = v
}

func minStepsBetween(cache *UpperTriangularMatrix, nums []int, l, r int) int {
	if l == r {
		return nums[l]
	}

	if v := cache.Get(l-1, r-1); v != 0 {
		return v
	}

	a, b := nums[l], minStepsBetween(cache, nums, l+1, r)
	if a > b {
		cache.Set(l-1, r-1, b)
		return b
	}
	cache.Set(l-1, r-1, a)
	return a
}

func jump(nums []int) int {
	numsLen := len(nums)
	cache := NewUpperTriangularMatrix(numsLen - 1)
	nums[numsLen-1] = 0
	for i := numsLen - 2; i > -1; i-- {
		if nums[i] == 0 {
			nums[i] = numsLen
			continue
		}

		l, r := i+1, i+nums[i]
		if r > numsLen-1 {
			r = numsLen - 1
		}
		nums[i] = minStepsBetween(cache, nums, l, r) + 1
	}

	return nums[0]

}

func minStepsBetween2(cache map[[2]int]int, nums []int, l, r int) int {
	if l == r {
		return nums[l]
	}

	key := [2]int{l, r}
	if v, ok := cache[key]; ok {
		return v
	}

	a, b := nums[l], minStepsBetween2(cache, nums, l+1, r)
	if a > b {
		cache[key] = b
		return b
	}
	return a
}

// 8 ms	6.6 MB
func jump2(nums []int) int {
	numsLen := len(nums)
	cache := make(map[[2]int]int)
	nums[numsLen-1] = 0
	for i := numsLen - 2; i > -1; i-- {
		if nums[i] == 0 {
			nums[i] = numsLen
			continue
		}

		l, r := i+1, i+nums[i]
		if r > numsLen-1 {
			r = numsLen - 1
		}
		nums[i] = minStepsBetween2(cache, nums, l, r) + 1
	}

	return nums[0]

}

func main() {
	fmt.Printf("r: %d\n", jump([]int{1, 1, 1}))
	fmt.Printf("r: %d\n", jump([]int{2, 1}))
	fmt.Printf("r: %d\n", jump([]int{2, 3, 0, 1, 4}))
	fmt.Printf("r: %d\n", jump([]int{2, 3, 1, 1, 4}))
}
