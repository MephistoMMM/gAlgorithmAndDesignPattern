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

// https://leetcode.com/problems/minimum-path-sum/

// 8 ms	3.9 MB
func minPathSumDynamically(grid [][]int, m, n int) int {
	// first column
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}

	// first row
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	// other cells
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i-1][j] > grid[i][j-1] {
				grid[i][j] += grid[i][j-1]
			} else {
				grid[i][j] += grid[i-1][j]
			}
		}
	}

	return grid[m-1][n-1]
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	return minPathSumDynamically(grid, m, n)
}

func main() {
	cnsl := &utils.Console{}
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	cnsl.Value(minPathSum(grid))
}
