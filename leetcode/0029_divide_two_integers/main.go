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

// https://leetcode.com/problems/divide-two-integers/

// solution: https://leetcode.com/problems/divide-two-integers/discuss/13420/32-times-bit-shift-operation-in-C-with-O(1)-solution

func divide(dividend int, divisor int) int {
	//special cases
	if divisor == 0 || (dividend == math.MinInt32 && divisor == -1) {
		return math.MaxInt32
	}

	// transform to unsigned int
	sign := (dividend > 0) == (divisor > 0)
	A := divisor
	if divisor < 0 {
		A = -divisor
	}
	B := dividend
	if dividend < 0 {
		B = -dividend
	}
	ret := 0

	// shift 32 times
	for i := 32; i > 0; i-- {
		k := uint(i - 1)
		if (B >> k) >= A {
			ret = (ret << 1) | 0x01
			B -= (A << k) // update B
		} else {
			ret = ret << 1
		}
	}

	if !sign {
		ret = -ret
	}

	return ret
}

func main() {
	fmt.Printf("%d \n", divide(10, 3))
	fmt.Printf("%d \n", divide(7, -3))
}
