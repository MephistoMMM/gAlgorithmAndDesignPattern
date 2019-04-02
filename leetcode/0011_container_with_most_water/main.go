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

// https://leetcode.com/problems/container-with-most-water/

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// A simple proof:

// case hi <= hj:
// we can prove that j is the best choice(within the range from i to j) for i
// for any k(i < k < j):
// area(i, j) == min(hi, hj)*(j-i) == hi*(j-i) > hi*(k-i) >= min(hi,hk)*(k-i) == area(i, k)

// then we have area(i, j) > area(i, k)
// it means j is the best choice for i, area(i,j)= largest_area_involves_i
// so, max_area_of_range(i, j) == max(max_area_of_range(i+1, j), largest_area_involves_i)
// case hi >= hj:
// similarly, we can prove that max_area_of_range(i, j) == max(max_area_of_range(i+1, j), area(i, j))
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0

	for i != j {
		volume := minInt(height[i], height[j]) * (j - i)

		if volume > max {
			max = volume
		}

		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return max
}

func main() {
	fmt.Printf("r: %d ", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
