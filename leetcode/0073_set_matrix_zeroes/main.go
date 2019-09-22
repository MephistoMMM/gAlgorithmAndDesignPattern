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
	"github.com/MephistoMMM/gAlgorithmAndDesignPattern/leetcode/utils"
	"math"
)

// https://leetcode.com/problems/set-matrix-zeroes/

type Status []uint64

func NewStatus(n int) Status {
	return make([]uint64, int(math.Ceil(float64(n)/64)))
}

// SetZero let n byte present zero
func (st Status) SetZero(n int) {
	index := n / 64
	offset := uint8(n % 64)

	st[index] |= uint64(1) << offset
}

func (st Status) IsZero(n int) bool {
	index := n / 64
	offset := uint8(n % 64)

	v := st[index] & (uint64(1) << offset)
	return 0 < v
}

// 16 ms	7.3 MB
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	rowStatus := NewStatus(m)
	columnStatus := NewStatus(n)

	// scan matrix to record zero status
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rowStatus.SetZero(i)
				columnStatus.SetZero(j)
			}
		}
	}

	// clean row and column whoes status is zero
	for i := 0; i < m; i++ {
		if !rowStatus.IsZero(i) {
			continue
		}

		for j := 0; j < n; j++ {
			matrix[i][j] = 0
		}
	}

	for j := 0; j < n; j++ {
		if !columnStatus.IsZero(j) {
			continue
		}

		for i := 0; i < m; i++ {
			matrix[i][j] = 0
		}
	}

}

func main() {
	cnsl := &utils.Console{}
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	matrix2 := [][]int{
		{3, 5, 5, 6, 9, 1, 4, 5, 0, 5},
		{2, 7, 9, 5, 9, 5, 4, 9, 6, 8},
		{6, 0, 7, 8, 1, 0, 1, 6, 8, 1},
		{7, 2, 6, 5, 8, 5, 6, 5, 0, 6},
		{2, 3, 3, 1, 0, 4, 6, 5, 3, 5},
		{5, 9, 7, 3, 8, 8, 5, 1, 4, 3},
		{2, 4, 7, 9, 9, 8, 4, 7, 3, 7},
		{3, 5, 2, 8, 8, 2, 2, 4, 9, 8},
	}
	setZeroes(matrix2)
	cnsl.DoubleDimArray(matrix2)
	setZeroes(matrix)
	cnsl.DoubleDimArray(matrix)
}
