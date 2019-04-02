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

// https://leetcode.com/problems/integer-to-roman/

var charlist = []struct {
	char byte
	val  int
}{
	{'M', 1000},
	{'D', 500},
	{'C', 100},
	{'L', 50},
	{'X', 10},
	{'V', 5},
	{'I', 1},
}

// 16 ms 3.3 MB
func intToRoman2(num int) string {
	result := make([]byte, 0, 16)
	for i := 0; num != 0; i++ {
		q := num / charlist[i].val
		num = num % charlist[i].val
		if q != 0 {
			for k := 0; k < q; k++ {
				result = append(result, charlist[i].char)
			}
		}

		if i == len(charlist)-1 {
			break
		}

		subChar := charlist[i+(2-i%2)]
		q = num / (charlist[i].val - subChar.val)
		num = num % (charlist[i].val - subChar.val)
		if q != 0 {
			result = append(result, subChar.char, charlist[i].char)
		}

	}

	return string(result)
}

type RomanChar struct {
	char []byte
	val  int
}

var chars = []RomanChar{
	{[]byte("M"), 1000},
	{[]byte("CM"), 900},
	{[]byte("D"), 500},
	{[]byte("CD"), 400},
	{[]byte("C"), 100},
	{[]byte("XC"), 90},
	{[]byte("L"), 50},
	{[]byte("XL"), 40},
	{[]byte("X"), 10},
	{[]byte("IX"), 9},
	{[]byte("V"), 5},
	{[]byte("IV"), 4},
	{[]byte("I"), 1},
}

// 16 ms	3.3 MB
func intToRoman(num int) string {
	result := make([]byte, 0, 16)
	for i := 0; num != 0; i++ {
		q := num / chars[i].val
		num = num % chars[i].val
		if q == 0 {
			continue
		}

		for k := 0; k < q; k++ {
			result = append(result, chars[i].char...)
		}

	}

	return string(result)
}

func main() {
	fmt.Printf("r: %s ", intToRoman(1994))
	fmt.Printf("r: %s ", intToRoman2(1994))
}
