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

// https://leetcode.com/problems/rotate-image/

func mapClockwise90Degrees(row, i, j int) (ni, nj int) {
	ni = j
	nj = row - i - 1
	return
}

func rotatePosition(matrix [][]int, i, j int) {
	tmp := matrix[i][j]
	i, j = mapClockwise90Degrees(len(matrix), i, j)
	matrix[i][j], tmp = tmp, matrix[i][j]
	i, j = mapClockwise90Degrees(len(matrix), i, j)
	matrix[i][j], tmp = tmp, matrix[i][j]
	i, j = mapClockwise90Degrees(len(matrix), i, j)
	matrix[i][j], tmp = tmp, matrix[i][j]
	i, j = mapClockwise90Degrees(len(matrix), i, j)
	matrix[i][j], tmp = tmp, matrix[i][j]
}

// 0 ms	2.6 MB
func rotate(matrix [][]int) {
	row := len(matrix)
	for i := 0; i < row/2; i++ {
		for j, ej := i, row-i-1; j < ej; j++ {
			rotatePosition(matrix, i, j)
		}
	}
}

func showMatrix(matrix [][]int) {
	fmt.Println("----------------------------")
	for _, v := range matrix {
		fmt.Printf("%v\n", v)
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	showMatrix(matrix)
	matrix2 := [][]int{
		{1, 2, 3, 10},
		{4, 5, 6, 11},
		{7, 8, 9, 12},
		{13, 14, 15, 16},
	}
	rotate(matrix2)
	showMatrix(matrix2)
}
