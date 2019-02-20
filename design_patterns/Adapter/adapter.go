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
package adapter

import "fmt"

// Banner is Adaptee
type Banner struct {
	str string
}

func NewBanner(str string) Banner {
	return Banner{str: str}
}

// ShowWithParen ...
func (br Banner) ShowWithParen() {
	fmt.Printf("(%s)", br.str)
}

// ShowWithAster ...
func (br Banner) ShowWithAster() {
	fmt.Printf("*%s*", br.str)
}

// Print is Target
type Print interface {
	PrintWeak()
	PrintStrong()
}

// PrintBanner is Adapter
type PrintBanner struct {
	Banner
}

func NewPrintBanner(str string) PrintBanner {
	return PrintBanner{
		Banner: NewBanner(str),
	}
}

// PrintWeak ...
func (pb PrintBanner) PrintWeak() {
	pb.ShowWithParen()
}

// PrintStrong ...
func (pb PrintBanner) PrintStrong() {
	pb.ShowWithAster()
}

func examplePrint(p Print) {
	p.PrintWeak()
	p.PrintStrong()
}

func exampleMain() {
	pb := NewPrintBanner("sdafsdf")
	examplePrint(pb)
}
