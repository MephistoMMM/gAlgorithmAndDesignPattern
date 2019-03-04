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

import "sync"

type PrinterProxy struct {
	name string
	real *Printer

	mutex sync.Mutex
}

func NewPrinterProxy(name string) *PrinterProxy {
	return &PrinterProxy{name: name}
}

// SetPrinterName ...
func (pp *PrinterProxy) SetPrinterName(name string) {
	pp.mutex.Lock()
	defer pp.mutex.Unlock()

	if pp.real != nil {
		pp.real.SetPrinterName(name)
	}
	pp.name = name
}

// GetPrinterName ...
func (pp *PrinterProxy) GetPrinterName() string {
	return pp.name
}

// Print ...
func (pp *PrinterProxy) Print(str string) {
	pp.realize()
	pp.real.Print(str)
}

// realize ...
func (pp *PrinterProxy) realize() {
	pp.mutex.Lock()
	defer pp.mutex.Unlock()

	if pp.real == nil {
		pp.real = NewPrinter(pp.name)
	}
}
