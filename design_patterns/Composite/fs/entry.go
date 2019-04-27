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
package fs

import (
	"errors"
	"fmt"
)

var (
	ErrFileTreatment = errors.New("ErrFileTreatment")
)

type Entry interface {
	Name() string
	Size() int
	Add(entry Entry) Entry
	PrintList(prefix string)
	String() string
}

type BaseEntry struct {
	name string
	size int
}

// Name ...
func (be *BaseEntry) Name() string {
	return be.name
}

// Size ...
func (be *BaseEntry) Size() int {
	return be.size
}

// Add ...
func (be *BaseEntry) Add(entry Entry) Entry {
	panic(ErrFileTreatment)
}

// String ...
func (be *BaseEntry) String() string {
	return fmt.Sprintf("%s (%d)", be.Name(), be.Size())
}

func PrintList(entry Entry) {
	entry.PrintList("")
}
