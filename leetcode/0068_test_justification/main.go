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

import "gAaD/leetcode/utils"

// https://leetcode.com/problems/text-justification/

type Line struct {
	size int

	words       []string
	left, right int
	chars       int
}

func NewLine(words []string, size, start int) *Line {
	if start >= len(words) {
		return nil
	}
	return &Line{size, words, start, start + 1, len(words[start])}
}

// Expand ...
func (line *Line) Expand() int {
	charsNum := line.right - line.left + line.chars - 1
	for line.right < len(line.words) {
		charsNum += 1 + len(line.words[line.right])
		if charsNum > line.size {
			break
		}

		line.chars += len(line.words[line.right])
		line.right += 1
	}
	return line.right
}

func (line *Line) finalString() string {
	byts := make([]byte, 0, line.size)
	writedSize := 0
	for _, word := range line.words[line.left:line.right] {
		// write word
		byts = append(byts, []byte(word)...)
		writedSize += len(word)
		if writedSize != line.size {
			byts = append(byts, ' ')
			writedSize++
		}
	}

	for writedSize != line.size {
		byts = append(byts, ' ')
		writedSize += 1
	}

	return string(byts)
}

func (line *Line) normalString() string {
	num := line.right - line.left - 1
	spaces := line.size - line.chars

	var avg, reminder int
	if num == 0 {
		avg, reminder = spaces, 0
	} else {
		avg, reminder = spaces/num, spaces%num
	}

	byts := make([]byte, 0, line.size)
	writedSize := 0
	for _, word := range line.words[line.left:line.right] {
		// write word
		byts = append(byts, []byte(word)...)
		writedSize += len(word)

		if writedSize == line.size {
			break
		}

		// write space
		spacesNum := avg
		if reminder > 0 {
			spacesNum++
			reminder--
		}
		writedSize += spacesNum
		for spacesNum > 0 {
			byts = append(byts, ' ')
			spacesNum--
		}
	}

	return string(byts)
}

func (line *Line) String(isFinal bool) string {
	if isFinal {
		return line.finalString()
	}
	return line.normalString()
}

//0 ms	2.2 MB
func fullJustify(words []string, maxWidth int) []string {
	lines := []*Line{}

	start := 0
	var line *Line
	for start < len(words) {
		line = NewLine(words, maxWidth, start)
		start = line.Expand()
		lines = append(lines, line)
	}

	lineLen := len(lines)
	result := make([]string, lineLen)
	for i := range lines[0 : lineLen-1] {
		result[i] = lines[i].String(false)
	}
	result[lineLen-1] = lines[lineLen-1].String(true)

	return result
}

func main() {
	cnsl := &utils.Console{}
	// cnsl.List(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
	// cnsl.List(fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))
	cnsl.List(fullJustify([]string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"}, 16))
}
