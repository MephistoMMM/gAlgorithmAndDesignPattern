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

// https://leetcode.com/problems/unique-paths-ii/

// 0 ms	2.6 MB
func uniquePathsDynamically(obstacleGrid [][]int, m, n int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	// first column
	isBlocked := false
	for i := range obstacleGrid {
		if isBlocked || obstacleGrid[i][0] == 1 {
			obstacleGrid[i][0] = 0
			isBlocked = true
		} else {
			obstacleGrid[i][0] = 1
		}
	}

	// first row
	obstacleGrid[0][0] = 0
	isBlocked = false
	for j := range obstacleGrid[0] {
		if isBlocked || obstacleGrid[0][j] == 1 {
			obstacleGrid[0][j] = 0
			isBlocked = true
		} else {
			obstacleGrid[0][j] = 1
		}
	}

	// other cells
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}

			obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
		}
	}

	return obstacleGrid[m-1][n-1]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	return uniquePathsDynamically(obstacleGrid, m, n)
}

func main() {
	cnsl := &utils.Console{}
	obstacle1 := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	obstacle2 := [][]int{
		{1, 1},
	}
	obstacle3 := [][]int{
		{1},
	}
	obstacle4 := [][]int{
		{0, 0},
		{1, 1},
		{0, 0},
	}
	obstacle5 := [][]int{
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	}
	cnsl.Value(uniquePathsWithObstacles(obstacle1))
	cnsl.Value(uniquePathsWithObstacles(obstacle2))
	cnsl.Value(uniquePathsWithObstacles(obstacle3))
	cnsl.Value(uniquePathsWithObstacles(obstacle4))
	cnsl.Value(uniquePathsWithObstacles(obstacle5))
}
