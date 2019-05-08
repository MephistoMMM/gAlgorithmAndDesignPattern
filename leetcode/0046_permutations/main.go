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

// https://leetcode.com/problems/permutations/

func swap(nums []int, i, j int) {
	if i == j {
		return
	}

	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func factorial(n int) int {
	factVal := 1
	if n < 0 {
		panic("Factorial of negative number doesn't exist.")
	} else {
		for i := 2; i <= n; i++ {
			factVal *= i
		}

	}
	return factVal
}

func permuteRecursively(result *[][]int, nums []int, i int) {
	if i == len(nums) {
		r := make([]int, len(nums))
		copy(r, nums)
		*result = append(*result, r)
		return
	}

	for j := i; j < len(nums); j++ {
		swap(nums, i, j)
		permuteRecursively(result, nums, i+1)
		swap(nums, i, j)
	}
}

// 4 ms	7.2 MB
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
				swap(nums, r, record[r]-1)
				continue
			} else {
				break
			}
		}

		swap(nums, r, record[r])
		record[r]++
		r++
	}
}

func permute(nums []int) [][]int {
	result := make([][]int, 0, factorial(len(nums)))
	permuteLooply(&result, nums)
	return result
}

// 4 ms	7.3 MB
func permute2(nums []int) [][]int {
	result := make([][]int, 0, factorial(len(nums)))
	permuteRecursively(&result, nums, 0)
	return result
}

func main() {
	fmt.Printf("%v \n", permute([]int{1, 2, 3}))
	fmt.Printf("%v \n", permute([]int{5, 4, 2, 6}))
}
