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

// https://leetcode.com/problems/spiral-matrix-ii/

// 0 ms	2.5 MB
func generateMatrix(n int) [][]int {
	row, column := n, n
	result := make([][]int, row)
	for i := range result {
		result[i] = make([]int, column)
	}

	bounds := [4]int{column, row, -1, 0}
	deltas := [4]int{1, 1, -1, -1}

	i, j := 0, -1
	cursors := [4]*int{&j, &i, &j, &i}
	// number of changing direction
	count := 0
	direction := 0
	cursor := cursors[0]
	for {

		index := direction % 4
		if *cursor+deltas[index] != bounds[index] {
			*cursor += deltas[index]
			count++
			result[i][j] = count
			continue
		}

		// cursor is at the bound
		bounds[index] -= deltas[index]
		direction++
		index = direction % 4
		cursor = cursors[index]

		// after modify curser
		if *cursor+deltas[index] == bounds[index] {
			break
		}
	}

	return result
}

func main() {
	cnsl := &utils.Console{}
	cnsl.DoubleDimArray(generateMatrix(3))
}
