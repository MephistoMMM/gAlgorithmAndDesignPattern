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

// https://leetcode.com/problems/permutations-ii/

func permuteLooply(result *[][]int, nums []int) {
	numsLen := len(nums)
	record := make([]int, numsLen)
	for i := 0; i < numsLen; i++ {
		record[i] = i
	}

	r := 0
	for {
		if r == numsLen {
			n := make([]int, numsLen)
			copy(n, nums)
			*result = append(*result, n)
			r--
		}

		if record[r] == numsLen {
			record[r] = r
			r--
			if r >= 0 {
				continue
			} else {
				break
			}
		}

		fmt.Printf("record: %v , nums: %v , r: %d \n", record, nums, r)
		for record[r]+1 < numsLen && nums[record[r]+1] == nums[record[r]] {
			record[r]++
		}

		if record[r] < numsLen {
		}
		record[r]++
		r++
	}
}

func permuteRecursively(result *[][]int, nums []int, used []bool, cur []int, i int) {
	numsLen := len(nums)

	if i == numsLen {
		n := make([]int, numsLen)
		copy(n, cur)
		*result = append(*result, n)
	}

	for k := 0; k < numsLen; k++ {
		if used[k] {
			continue
		}

		if k+1 < numsLen && nums[k] == nums[k+1] && !used[k+1] {
			continue
		}

		cur[i] = nums[k]
		used[k] = true
		permuteRecursively(result, nums, used, cur, i+1)
		used[k] = false
	}
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	cur := make([]int, len(nums))
	used := make([]bool, len(nums))
	permuteRecursively(&result, nums, used, cur, 0)
	return result
}

func main() {
	fmt.Printf("%v \n", permuteUnique([]int{0, 1, 0, 0, 9}))
	// fmt.Printf("%v \n", permuteUnique([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 1, 1, 1, 1, 1}))
	// fmt.Printf("%v \n", permuteUnique([]int{5, 4, 2, 6}))
}
