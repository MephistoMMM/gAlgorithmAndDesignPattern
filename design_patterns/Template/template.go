// Copyright © 2018 Mephis Pheies <mephistommm@gmail.com>
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
package template

import "fmt"

type AbstractDisplay interface {
	Open()
	Print()
	Close()
	Display()
}

type BaseDisplay struct{}

// Display ...
func (bd BaseDisplay) Display(display AbstractDisplay) {
	display.Open()
	for i := 0; i < 5; i++ {
		display.Print()
	}
	display.Close()
}

type CharDisplay struct {
	BaseDisplay
	char rune
}

func NewCharDisplay(char rune) *CharDisplay {
	return &CharDisplay{
		char: char,
	}
}

// Open ...
func (cd *CharDisplay) Open() {
	fmt.Print("<<")
}

// Print ...
func (cd *CharDisplay) Print() {
	fmt.Printf("%c", cd.char)
}

// Close ...
func (cd *CharDisplay) Close() {
	fmt.Println(">>")
}

func (cd *CharDisplay) Display() {
	cd.BaseDisplay.Display(cd)
}

type StringDisplay struct {
	BaseDisplay

	str   string
	width int
}

func NewStringDisplay(str string) *StringDisplay {
	return &StringDisplay{
		str:   str,
		width: len(str),
	}
}

// printLine ...
func (sd *StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < sd.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

// Open ...
func (sd *StringDisplay) Open() {
	sd.printLine()
}

// Print ...
func (sd *StringDisplay) Print() {
	fmt.Println("|" + sd.str + "|")
}

// Close ...
func (sd *StringDisplay) Close() {
	sd.printLine()
}

func (sd *StringDisplay) Display() {
	sd.BaseDisplay.Display(sd)
}

// 在模板方法没有多少成员时，直接用模板函数反而更加简单
