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

// https://leetcode.com/problems/string-to-integer-atoi/

func leftTrimSpace(s []byte) (r []byte) {
	validStart := len(s)
	for i, v := range s {
		if v != ' ' {
			validStart = i
			break
		}
	}

	return s[validStart:]
}

func isNumber(b byte) bool {
	return '0' <= b && b <= '9'
}

func isSign(b byte) bool {
	return b == '+' || b == '-'
}

func myAtoi(str string) int {
	x := leftTrimSpace([]byte(str))

	if len(x) == 0 || !isNumber(x[0]) && !isSign(x[0]) {
		return 0
	}

	sign := 1
	if isSign(x[0]) {
		if x[0] == '-' {
			sign = -1
		}
		x = x[1:]
	}

	r := 0
	for _, v := range x {
		if !isNumber(v) {
			break
		}
		vn := int(v - '0')

		if sign > 0 && (r > math.MaxInt32/10 || (r == math.MaxInt32/10 && vn > 7)) {
			return math.MaxInt32
		}

		if sign < 0 && (r < math.MinInt32/10 || (r == math.MinInt32/10 && vn > 8)) {
			return math.MinInt32
		}

		r = r*10 + vn*sign
	}

	return r
}

func main() {
	fmt.Printf("r: %d ", myAtoi("  -4193 with words"))
	fmt.Printf("r: %d ", myAtoi("-2147483649"))
	fmt.Printf("r: %d ", myAtoi("  "))
	fmt.Printf("r: %d ", myAtoi(" 2147483800"))
}
