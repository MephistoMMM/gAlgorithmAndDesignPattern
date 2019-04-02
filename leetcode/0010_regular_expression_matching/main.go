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
	"regexp"
)

// https://leetcode.com/problems/regular-expression-matching/

type VarBool struct {
	v bool
}

func createMemo(sLen, pLen int) [][]*VarBool {
	memo := make([][]*VarBool, sLen+1)
	for i := range memo {
		memo[i] = make([]*VarBool, pLen+1)
	}

	return memo
}

func isSame(sChar byte, pChar byte) bool {
	if pChar == '.' || pChar == sChar {
		return true
	}

	return false
}

var re = regexp.MustCompile(`(.\*)`)

// 8 ms 6.4 MB
func isMatch(s, p string) bool {
	reP := regexp.MustCompile("^" + re.ReplaceAllString(p, "($1)*") + "$")
	return reP.MatchString(s)
}

// 4 ms	2.8 MB
func isMatchDp2(s, p string) bool {
	memo := createMemo(len(s), len(p))

	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		if memo[i][j] == nil {
			ans := false
			if j == len(p) {
				ans = i == len(s)
			} else {
				currentMatch := i < len(s) && isSame(s[i], p[j])
				if j+1 < len(p) && p[j+1] == '*' {
					ans = dp(i, j+2) || currentMatch && dp(i+1, j)
				} else {
					ans = currentMatch && dp(i+1, j+1)
				}
			}

			memo[i][j] = &VarBool{ans}
		}

		return memo[i][j].v
	}

	return dp(0, 0)
}

// 4 ms	  3.2 MB
func isMatchDp(s, p string) bool {
	memo := make(map[[2]int]bool)

	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		key := [2]int{i, j}
		if _, ok := memo[key]; !ok {
			ans := false
			if j == len(p) {
				ans = i == len(s)
			} else {
				currentMatch := i < len(s) && isSame(s[i], p[j])
				if j+1 < len(p) && p[j+1] == '*' {
					ans = dp(i, j+2) || currentMatch && dp(i+1, j)
				} else {
					ans = currentMatch && dp(i+1, j+1)
				}
			}

			memo[key] = ans
		}

		return memo[key]
	}

	return dp(0, 0)
}

// 20 ms	2.2 MB
func isMatchRSC(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	firstMatch := len(s) > 0 && isSame(s[0], p[0])
	if len(p) >= 2 && p[1] == '*' {
		// consume x* or not
		return isMatchRSC(s, p[2:]) || firstMatch && isMatchRSC(s[1:], p)
	}

	return firstMatch && isMatchRSC(s[1:], p[1:])
}

func isMatchLegacy(s string, p string) bool {
	i, j := 0, 0
	prePChar := byte(0)
	for i < len(s) && j < len(p) {
		if p[j] == '*' {
			if !isSame(s[i], prePChar) {
				prePChar = '*'
				j++
				continue
			}

			// s not consume
			if isMatchLegacy(s[i:], p[j+1:]) {
				return true
			}

			// s consume
			i++
			continue
		}

		// try consume x*
		if j+1 < len(p) && p[j+1] == '*' && isMatchLegacy(s[i:], p[j+2:]) {
			return true
		}

		// same
		if isSame(s[i], p[j]) {
			prePChar = p[j]
			i++
			j++
		} else {
			return false
		}
	}

	if j < len(p) && p[j] == '*' {
		j++
	}

	// p has not been consumed
	for j+1 < len(p) && p[j+1] == '*' {
		j += 2
	}

	// True only both s and p are consumed
	if i == len(s) && j == len(p) {
		return true
	}

	return false
}

func main() {

	fmt.Printf("3: %t ", isMatch("bbbba", ".*a*a"))
	fmt.Printf("3: %t ", isMatch("abc", ".*"))
	fmt.Printf("1: %t ", isMatch("mississippi", "mis*is*p*."))
	fmt.Printf("2: %t ", isMatch("aab", "ca*b"))
	fmt.Printf("2: %t ", isMatch("aaa", "a*a"))
	fmt.Printf("2: %t ", isMatch("a", "ab*"))
}
