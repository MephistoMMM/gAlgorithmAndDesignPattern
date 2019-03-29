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
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minInt(a, b int) int {
	if a > b {
		return b
	}

	return a
}

//           left_part          |        right_part
//   A[0], A[1], ..., A[i-1]  |  A[i], A[i+1], ..., A[m-1]
//   B[0], B[1], ..., B[j-1]  |  B[j], B[j+1], ..., B[n-1]
//
//   n+m is odd : max(A[i-1], B[j-1])
//   n+m is even : (max(A[i-1], B[j-1]) + min(A[i], B[j])) / 2
func findMedianSortedArraysInner(A []int, B []int) float64 {
	m := len(A)
	n := len(B)

	if m == 0 && n == 0 {
		// error
		return 0
	}
	if m == 0 {
		A = append([]int{B[0]}, B[n-1])
		m += 2
	}

	mini, maxi := 0, m

	// first find the correct i and j
	i, j := 0, 0
	for {
		i = (mini + maxi) / 2
		j = (m+n+1)/2 - i

		if (i == 0 || j == n || A[i-1] <= B[j]) && (i == m || j == 0 || A[i] >= B[j-1]) {
			break
		}

		if i > 0 && A[i-1] > B[j] {
			maxi = i
		}

		if j > 0 && B[j-1] > A[i] {
			mini = i + 1
		}
	}

	// second calculate median
	leftMax, rightMin := 0, 0
	if i == 0 {
		leftMax = B[j-1]
	}
	if j == 0 {
		leftMax = A[i-1]
	}
	if i == m {
		rightMin = B[j]
	}
	if j == n {
		rightMin = A[i]
	}

	if i > 0 && j > 0 {
		leftMax = maxInt(A[i-1], B[j-1])
	}
	if i < m && j < n {
		rightMin = minInt(A[i], B[j])
	}

	if (m+n)%2 == 1 {
		return float64(leftMax)
	}

	return float64(leftMax+rightMin) / 2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArraysInner(nums2, nums1)
	}

	return findMedianSortedArraysInner(nums1, nums2)
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Printf("r: %f.", findMedianSortedArrays(nums1, nums2))
}
