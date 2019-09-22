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
	"strings"
)

// https://leetcode.com/problems/valid-number/

func isUnsignedInt(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}

	return true
}

func isInt(s string) bool {
	if len(s) == 0 {
		return false
	}

	// has sign
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}

	return isUnsignedInt(s)
}

func isFloat(s string) bool {
	if len(s) == 0 {
		return false
	}

	// has sign
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}

	numbers := strings.Split(s, ".")
	if len(numbers) == 1 {
		return isUnsignedInt(numbers[0])
	}
	if len(numbers) > 2 {
		return false
	}

	if len(numbers[0]) == 0 {
		return isUnsignedInt(numbers[1])
	}

	if len(numbers[1]) == 0 {
		return isUnsignedInt(numbers[0])
	}

	return isUnsignedInt(numbers[0]) && isUnsignedInt(numbers[1])
}

// 0 ms	2.3 MB
func isNumber(s string) bool {
	s = strings.TrimSpace(s)

	numbers := strings.Split(s, "e")
	if len(numbers) == 1 {
		return isFloat(numbers[0])
	}
	if len(numbers) == 2 {
		return isFloat(numbers[0]) && isInt(numbers[1])
	}

	return false
}

func main() {
	cnsl := &utils.Console{}
	cnsl.Value(isNumber("0"))
	cnsl.Value(isNumber(" 0.1 "))
	cnsl.Value(isNumber("abc"))
	cnsl.Value(isNumber(" -90e3   "))
	cnsl.Value(isNumber("e3"))
	cnsl.Value(isNumber("."))
	cnsl.Value(isNumber(".1"))
	cnsl.Value(isNumber("+.1"))
	cnsl.Value(isNumber("-1."))
}
