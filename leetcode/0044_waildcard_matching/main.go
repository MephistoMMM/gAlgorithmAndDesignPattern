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

// https://leetcode.com/problems/wildcard-matching/
// http://yucoding.blogspot.com/2013/02/leetcode-question-123-wildcard-matching.html

func isMatch(s string, p string) bool {
	sLen, pLen := len(s), len(p)
	if pLen == 0 && sLen != 0 {
		return false
	}

	i, j := 0, 0
	star, ss := -1, 0
	for i < sLen {
		if j < pLen && (p[j] == '?' || s[i] == p[j]) {
			i++
			j++
			continue
		}
		if j < pLen && p[j] == '*' {
			star = j
			ss = i
			j++
			continue
		}
		if star >= 0 {
			j = star + 1
			ss++
			i = ss
			continue
		}

		return false
	}

	for j < pLen && p[j] == '*' {
		j++
	}

	return j == pLen
}

func main() {

	fmt.Printf("r: %t \n", isMatch("cb", "?a"))
	fmt.Printf("r: %t \n", isMatch("adceb", "*a*b"))
	fmt.Printf("r: %t \n", isMatch("acdcb", "a*c?b"))
	fmt.Printf("r: %t \n", isMatch("", "*"))
	fmt.Printf("r: %t \n", isMatch("aa", "*"))
	fmt.Printf("r: %t \n", isMatch("bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab", "b*b*ab**ba*b**b***bba"))
}
