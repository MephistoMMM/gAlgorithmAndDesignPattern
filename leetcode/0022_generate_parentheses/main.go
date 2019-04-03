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
	"math"
)

// https://leetcode.com/problems/generate-parentheses/

func genParenthByBackTrace(result *[]string, leftL int, leftR int, current []byte, index int) {
	if leftL == 0 && leftR == 0 {
		*result = append(*result, string(current))
	}

	if leftL > 0 {
		current[index] = '('
		genParenthByBackTrace(result, leftL-1, leftR+1, current, index+1)
	}

	if leftR > 0 {
		current[index] = ')'
		genParenthByBackTrace(result, leftL, leftR-1, current, index+1)
	}

}

// 8 ms	7.3 MB
func generateParenthesis(n int) []string {
	// An = (3**(n-1) + 1)/2
	result := make([]string, 0, (int(math.Pow(3, float64(n-1)))+1)/2)
	current := make([]byte, n*2)

	genParenthByBackTrace(&result, n, 0, current, 0)
	return result
}

func main() {
	fmt.Printf("r: %v \n", generateParenthesis(4))
}
