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

// https://leetcode.com/problems/search-a-2d-matrix/

type SortedIntArray interface {
	Len() int
	Get(i int) int
}

type Matrix [][]int

func (m Matrix) Len() int {
	return len(m)
}

func (m Matrix) Get(i int) int {
	return m[i][0]
}

type MatrixRow []int

func (mr MatrixRow) Len() int {
	return len(mr)
}

func (mr MatrixRow) Get(i int) int {
	return mr[i]
}

func binarySearch(arr SortedIntArray, aim int) int {
	l, r := 0, arr.Len()-1
	for l <= r {
		m := (l + r) / 2
		mv := arr.Get(m)
		if aim == mv {
			return m
		}
		if aim < mv {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return r
}

// 4 ms	3.8 MB
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m := Matrix(matrix)
	r := binarySearch(m, target)
	if r < 0 {
		return false
	}

	row := MatrixRow(matrix[r])
	i := binarySearch(row, target)
	if row[i] != target {
		return false
	}
	return true
}

func main() {
	cnsl := &utils.Console{}
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	cnsl.Value(searchMatrix(matrix, 3))
	matrix2 := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	cnsl.Value(searchMatrix(matrix2, 13))
}
