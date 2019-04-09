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

// https://leetcode.com/problems/substring-with-concatenation-of-all-words/
func initWordsDict(words []string) map[string]int {
	dict := make(map[string]int)

	for _, v := range words {
		dict[v]++
	}

	return dict
}

func copyWordsDict(dst map[string]int, src map[string]int) map[string]int {
	if dst == nil {
		dst = make(map[string]int)
	}

	for k, v := range src {
		dst[k] = v
	}

	return dst
}

// 36 ms 6 MB
func findSubstring(s string, words []string) []int {
	result := []int{}
	if words == nil || len(words) == 0 {
		return result
	}

	dict := initWordsDict(words)

	sl, wl := len(s), len(words[0])
	ws := len(words)
	tdict := copyWordsDict(nil, dict)
	for i := 0; i <= sl-ws*wl; i++ {
		p := i
		if dict[s[p:p+wl]] == 0 {
			continue
		}

		tdict = copyWordsDict(nil, dict)
		valid := true
		for j := p; j < p+wl*ws; j += wl {
			if tdict[s[j:j+wl]] == 0 {
				valid = false
				break
			}

			tdict[s[j:j+wl]]--
		}

		if !valid {
			continue
		}

		result = append(result, p)
	}

	return result
}

func main() {
	fmt.Printf("%v \n", findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Printf("%v \n", findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}
