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

// https://leetcode.com/problems/unique-paths/

// time limit exceeded
func uniquePathsRecursively(m int, n int, i int, j int, counter *int) {
	if i == m && j == n {
		*counter += 1
		return
	}

	if i < m {
		// go down
		uniquePathsRecursively(m, n, i+1, j, counter)
	}

	if j < n {
		// go right
		uniquePathsRecursively(m, n, i, j+1, counter)
	}
}

func uniquePaths2(m int, n int) int {
	counter := 0
	uniquePathsRecursively(m-1, n-1, 0, 0, &counter)
	return counter
}

func factorial(from, to int) int {
	result := 1
	for from > to {
		result *= from
		from--
	}
	return result
}

func Select(num, total int) int {
	if num > total/2 {
		num = total - num
	}
	if num == 0 {
		return 1
	}
	return factorial(total, total-num) / factorial(num, 1)
}

// 0 ms	1.9 MB
func uniquePaths(m int, n int) int {
	return Select(n-1, m+n-2)
}

func main() {
	cnsl := &utils.Console{}
	cnsl.Value(uniquePaths(3, 2))
	cnsl.Value(uniquePaths(7, 3))
	cnsl.Value(uniquePaths(1, 100))
}
