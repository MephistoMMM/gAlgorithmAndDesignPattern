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

import "gAaD/leetcode/utils"

// https://leetcode.com/problems/n-queens-ii/

// 0 ms	1.9 MB
func solveNQueensRecursively(n int,
	i int, state []bool, pos []int, results *int) {
	// finish putting Queens
	if i == n {
		*results += 1
		return
	}

	for j, v := range state {
		// the vth position has been hold
		if v {
			continue
		}

		// check safety at access cornor
		safe := true
		for k := 0; k < i; k++ {
			if i-k == j-pos[k] || i-k == pos[k]-j {
				safe = false
				break
			}
		}

		if !safe {
			continue
		}

		// the vth position is empty
		pos[i] = j
		state[j] = true

		solveNQueensRecursively(n,
			i+1, state, pos, results)
		state[j] = false
	}

}

func totalNQueens(n int) int {
	results := 0
	if n <= 0 {
		return 0
	}

	state := make([]bool, n)
	pos := make([]int, n)
	solveNQueensRecursively(n, 0, state, pos, &results)
	return results
}

func main() {
	cnsl := &utils.Console{}
	cnsl.Valuef("%d: %d \n ", 4, totalNQueens(4))
	cnsl.Valuef("%d: %d \n ", 5, totalNQueens(5))
	cnsl.Valuef("%d: %d \n ", 6, totalNQueens(6))
	cnsl.Valuef("%d: %d \n ", 7, totalNQueens(7))
	cnsl.Valuef("%d: %d \n ", 8, totalNQueens(8))
	cnsl.Valuef("%d: %d \n ", 9, totalNQueens(9))
	cnsl.Valuef("%d: %d \n ", 10, totalNQueens(10))

}
