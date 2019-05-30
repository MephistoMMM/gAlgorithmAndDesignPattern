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

// https://leetcode.com/problems/powx-n/

// 0 ms	2.1 MB
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	reset := 1.0
	for n != 1 {
		mod := n % 2
		if mod == 0 {
			x *= x
		} else {
			reset *= x
			x *= x
		}

		n /= 2
	}

	return x * reset
}

// 0 ms	2.1 MB
func myPow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	mod := n % 2
	if mod == 0 {
		return myPow(x*x, n/2)
	}

	return myPow(x*x, n/2) * x
}

func main() {
	fmt.Printf("%f\n", myPow(2.10000, 3))
	fmt.Printf("%f\n", myPow(2.00000, 10))
	fmt.Printf("%f\n", myPow(2.00000, -2))
}
