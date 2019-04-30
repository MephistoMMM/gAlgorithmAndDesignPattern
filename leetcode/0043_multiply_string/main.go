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

// https://leetcode.com/problems/multiply-strings/
// may right https://leetcode.com/problems/multiply-strings/discuss/17771/Very-concise-16-ms-c%2B%2B-solution

import "fmt"

func reverseBytes(byts []byte) {
	for i, j := 0, len(byts)-1; i < j; i, j = i+1, j-1 {
		byts[i], byts[j] = byts[j], byts[i]
	}
}

func digitMultiplyWithTenPower(accu []byte, num string, digit byte, power int) []byte {
	digit -= '0'

	accu = accu[:0]
	for power > 0 {
		accu = append(accu, '0')
		power--
	}

	carry := byte(0)
	for i := len(num) - 1; i > -1; i-- {
		v := byte(num[i] - '0')
		r := v*digit + carry
		carry = r / 10
		r = r % 10
		accu = append(accu, r+'0')
	}

	if carry > 0 {
		accu = append(accu, carry+'0')
	}

	return accu
}

// add b to a
func Add(sum []byte, v []byte) []byte {
	a := sum
	sum = sum[:0]
	carry := byte(0)

	i := 0
	for ; i < len(a) && i < len(v); i++ {
		r := (a[i] - '0') + (v[i] - '0') + carry
		carry = r / 10
		r = r % 10
		sum = append(sum, r+'0')

	}
	for i < len(a) {
		r := a[i] - '0' + carry
		carry = r / 10
		r = r % 10
		sum = append(sum, r+'0')
		i++
	}
	for i < len(v) {
		r := v[i] - '0' + carry
		carry = r / 10
		r = r % 10
		sum = append(sum, r+'0')
		i++
	}
	if carry > 0 {
		sum = append(sum, carry+'0')
	}

	return sum
}

// 4 ms	2.2 MB
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	num1Len, num2Len := len(num1), len(num2)
	sum := make([]byte, 0, num1Len+num2Len)
	accu := make([]byte, 0, num1Len+num2Len)
	for i := 0; i < num2Len; i++ {
		sum = Add(sum, digitMultiplyWithTenPower(accu, num1, byte(num2[i]), num2Len-i-1))
	}

	reverseBytes(sum)
	return string(sum)
}

func main() {
	fmt.Printf("r: %s \n", multiply("123", "456"))
}
