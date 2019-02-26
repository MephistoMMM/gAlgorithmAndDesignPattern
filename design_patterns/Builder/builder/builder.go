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
package builder

import (
	"bufio"
	"bytes"
	"os"
)

// Builder interface in Builder design pattern
type Builder interface {
	// MakeTitle ...
	MakeTitle(title string)
	// MakeString ...
	MakeString(str string)
	// Make Iterm ...
	MakeItems(items []string)
	// Close ...
	Close()
}

// TextBuilder implements Builder to provide functions writing document
// to normal text.
type TextBuilder struct {
	// buffer save document content
	buffer bytes.Buffer
}

func NewTextBuilder() *TextBuilder {
	return &TextBuilder{}
}

// MakeTitle  ...
func (tb *TextBuilder) MakeTitle(title string) {
	tb.buffer.WriteString("============================\n")
	tb.buffer.WriteString("[" + title + "]\n")
	tb.buffer.WriteString("\n")
}

// MakeString  ...
func (tb *TextBuilder) MakeString(str string) {
	tb.buffer.WriteString("* " + str + "]\n")
	tb.buffer.WriteString("\n")
}

// MakeItems  ...
func (tb *TextBuilder) MakeItems(items []string) {
	for _, v := range items {
		tb.buffer.WriteString("  -" + v + "\n")
	}
	tb.buffer.WriteString("\n")
}

// Close  ...
func (tb *TextBuilder) Close() {
	tb.buffer.WriteString("============================\n")
}

func (tb *TextBuilder) Result() string {
	return tb.buffer.String()
}

type HTMLBuilder struct {
	filename string
	file     *os.File

	// wrap os.File
	writer *bufio.Writer
}

// HTMLBuilder implements Builder to provide functions writing document
// to html file.
func NewHTMLBuilder(filename string) *HTMLBuilder {
	return &HTMLBuilder{
		filename: filename,
	}
}

// MakeTitle  ...
func (hb *HTMLBuilder) MakeTitle(title string) {
	if hb.file == nil {
		file, err := os.OpenFile(hb.filename, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			panic(err)
		}
		hb.file = file
		hb.writer = bufio.NewWriter(hb.file)
	}

	hb.writer.WriteString("<html><head><title>\n")
	hb.writer.WriteString(title)
	hb.writer.WriteString("</title></head>\n<body>\n")
	hb.writer.WriteString("<h1>")
	hb.writer.WriteString(title)
	hb.writer.WriteString("</h1>\n")
}

// MakeString  ...
func (hb *HTMLBuilder) MakeString(str string) {
	hb.writer.WriteString("<p>")
	hb.writer.WriteString(str)
	hb.writer.WriteString("</p>\n")
}

// MakeItems  ...
func (hb *HTMLBuilder) MakeItems(items []string) {
	hb.writer.WriteString("<ul>")
	for _, v := range items {
		hb.writer.WriteString("<li>")
		hb.writer.WriteString(v)
		hb.writer.WriteString("</li>\n")
	}
	hb.writer.WriteString("</ul>\n")
}

// Close  ...
func (hb *HTMLBuilder) Close() {
	hb.writer.WriteString("</body>\n<html>\n")
	hb.writer.Flush()
	if err := hb.file.Close(); err != nil {
		panic(err)
	}
	hb.file = nil
	hb.writer = nil
}

func (hb *HTMLBuilder) Result() string {
	return hb.filename
}
