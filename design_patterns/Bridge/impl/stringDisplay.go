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
package impl

import "fmt"

// StringDisplay display string
type StringDisplay struct {
	str   string
	width int
}

func NewStringDisplay(str string) *StringDisplay {
	return &StringDisplay{
		str:   str,
		width: len(str),
	}
}

// RawOpen ...
func (sd *StringDisplay) RawOpen() {
	sd.printLine()
}

// RawPrint ...
func (sd *StringDisplay) RawPrint() {
	fmt.Printf("|%s|\n", sd.str)
}

// RawClose ...
func (sd *StringDisplay) RawClose() {
	sd.printLine()
}

// printLine ...
func (sd *StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < sd.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}