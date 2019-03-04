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
package printer

import (
	"fmt"
	"time"
)

// Printable interface provide constant for Printer and PrinterProxy
type Printable interface {
	// SetPrinterName set name
	SetPrinterName(name string)
	// GetPrinterName get name
	GetPrinterName() string
	// Print show text
	Print(str string)
}

// Printer reality
type Printer struct {
	name string
}

func NewPrinter(name string) *Printer {
	pt := &Printer{name}
	pt.heavyJob("Printer is generating(" + name + ")")
	return pt
}

// SetPrinterName ...
func (pt *Printer) SetPrinterName(name string) {
	pt.name = name
}

// GetPrinterName ...
func (pt *Printer) GetPrinterName() string {
	return pt.name
}

// Print ...
func (pt *Printer) Print(str string) {
	fmt.Println("===" + pt.name + "===")
	fmt.Println(str)
}

// heavyJob ...
func (pt *Printer) heavyJob(msg string) {
	fmt.Print(msg)
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println("end.")
}
