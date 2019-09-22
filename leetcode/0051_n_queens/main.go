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
	"github.com/MephistoMMM/gAlgorithmAndDesignPattern/leetcode/utils"
)

// https://leetcode.com/problems/n-queens/

// 4 ms	9.1 MB
func solveNQueensRecursively(n int, empty []byte,
	i int, state []bool, pos []int, current []string, results *[][]string) {
	// finish putting Queens
	if i == n {
		tmp := make([]string, n)
		copy(tmp, current)
		*results = append(*results, tmp)
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
		empty[j] = 'Q'
		current[i] = string(empty)
		empty[j] = '.'
		state[j] = true

		solveNQueensRecursively(n, empty,
			i+1, state, pos, current, results)
		state[j] = false
	}

}

func solveNQueens(n int) [][]string {
	results := make([][]string, 0)
	if n <= 0 {
		return results
	}

	state := make([]bool, n)
	pos := make([]int, n)
	current := make([]string, n)
	empty := make([]byte, n)
	for i := 0; i < n; i++ {
		empty[i] = '.'
	}

	solveNQueensRecursively(n, empty,
		0, state, pos, current, &results)
	return results
}

func printColumn(v interface{}) {
	table := v.([]string)
	fmt.Println("[")
	for _, r := range table {
		fmt.Printf("\t%v\n", r)
	}
	fmt.Println("]")
}

func main() {
	cnsl := &utils.Console{}
	cnsl.DoubleDimArrayWithItemFunc(printColumn, solveNQueens(4))
}
