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

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

var numMapList = [9][]byte{
	[]byte("abc"),  // 2
	[]byte("def"),  // 3
	[]byte("ghi"),  // 4
	[]byte("jkl"),  // 5
	[]byte("mno"),  // 6
	[]byte("pqrs"), // 7
	[]byte("tuv"),  // 8
	[]byte("wxyz"), // 9
}

// 0 ms	2.6 MB
func letterCombinationsRSC2(digits string, current []byte, i int, result *[]string) {
	if len(digits) == 0 {
		*result = append(*result, string(current[:i]))
		return
	}

	n := digits[0]
	if n < '2' && n > '9' {
		letterCombinationsRSC2(digits[1:], current, i, result)
		return
	}

	numStr := numMapList[n-'2']
	for _, v := range numStr {
		current[i] = v
		letterCombinationsRSC2(digits[1:], current, i+1, result)
	}
}

func letterCombinations2(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	current := make([]byte, len(digits))
	result := make([]string, 0)
	letterCombinationsRSC2(digits, current, 0, &result)
	return result
}

// 0 ms	2.7 MB
func letterCombinationsRSC(digits string, current []byte, i int) []string {
	if len(digits) == 0 {
		return []string{string(current[:i])}
	}

	n := digits[0]
	if n < '2' && n > '9' {
		return letterCombinationsRSC(digits[1:], current, i)
	}

	result := []string{}
	numStr := numMapList[n-'2']
	for _, v := range numStr {
		current[i] = v
		result = append(result, letterCombinationsRSC(digits[1:], current, i+1)...)
	}

	return result
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	current := make([]byte, len(digits))
	return letterCombinationsRSC(digits, current, 0)
}

func main() {
	fmt.Printf("r: %v \n", letterCombinations("23"))
}
