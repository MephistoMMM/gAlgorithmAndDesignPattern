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

// https://leetcode.com/problems/trapping-rain-water/

// 4 ms	2.8 MB
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	sum := 0

	l, r := 0, len(height)-1
	leftMax, rightMax := height[l], height[r]
	for l < r {
		vl, vr := height[l], height[r]
		if vl <= vr {
			if vl < leftMax {
				sum += leftMax - vl
			} else {
				leftMax = vl
			}
			l++
		} else {
			if vr < rightMax {
				sum += rightMax - vr
			} else {
				rightMax = vr
			}
			r--
		}
	}

	return sum
}

// 4 ms	2.9 MB
func trap3(height []int) int {
	sum, cur := 0, 0
	stack := make([]int, len(height))
	top := -1

	// 0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1
	for cur < len(height) {
		for top > -1 && height[cur] > height[stack[top]] {
			t := stack[top]
			top--
			if top < 0 {
				break
			}
			distance := cur - stack[top] - 1
			boundedHeight := Min(height[cur], height[stack[top]]) - height[t]
			sum += distance * boundedHeight
		}

		top++
		stack[top] = cur
		cur++
	}

	return sum
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calculate(height []int) int {
	length := len(height)
	minBoard := Min(height[0], height[length-1])
	sum := minBoard * length
	for _, v := range height {
		if v > minBoard {
			v = minBoard
		}
		sum -= v
	}
	return sum
}

// 8 ms	2.9 MB
func trap2(height []int) int {
	length := len(height)
	if length < 2 {
		return 0
	}

	peaks := make([]int, 0, length/2+1)
	// state: -1 down, 1 up
	state := 1
	for i := 1; i < length; i++ {
		// peek to bottom
		if state == 1 && height[i] < height[i-1] {
			state = -1
			peaks = append(peaks, i-1)
		}
		// bottom to peak
		if state == -1 && height[i] > height[i-1] {
			state = 1
		}
	}

	if state == 1 {
		peaks = append(peaks, length-1)
	}

	sum := 0
	peakLen := len(peaks)
	for i := 0; i < peakLen-1; {
		j := i + 1
		maxIdx := j
		for ; j < peakLen; j++ {
			if height[peaks[j]] >= height[peaks[i]] {
				maxIdx = j
				break
			}
			if height[peaks[j]] > height[peaks[maxIdx]] {
				maxIdx = j
			}
		}

		sum += calculate(height[peaks[i] : peaks[maxIdx]+1])
		i = maxIdx
	}

	return sum
}

func main() {
	fmt.Printf("r: %d\n", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Printf("r: %d\n", trap([]int{2, 8, 5, 5, 6, 1, 7, 4, 5}))
	fmt.Printf("r: %d\n", trap([]int{0, 0, 0, 0, 0, 0, 0, 0, 2, 1, 2, 1}))
	fmt.Printf("r: %d\n", trap([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	fmt.Printf("r: %d\n", trap([]int{2, 0, 2}))
	fmt.Printf("r: %d\n", trap([]int{5, 4, 1, 2}))
	fmt.Printf("r: %d\n", trap([]int{5, 2, 1, 2, 1, 5}))

}
