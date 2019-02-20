// * Copyright © 2018 Mephis Pheies <mephistommm@gmail.com>
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
// * Codes
package prototype

import "fmt"

type Cloneable interface {
	Clone() interface{}
}

// ** Product Interface

type Product interface {
	Cloneable

	Use(s string)
	CreateClone() Product
}

// ** Manager

type Manager struct {
	showcase map[string]Product
}

// Register ...
func (m *Manager) Register(name string, proto Product) {
	if m.showcase == nil {
		m.showcase = make(map[string]Product)
	}
	m.showcase[name] = proto
}

// Create ...
func (m *Manager) Create(protoname string) Product {
	p := m.showcase[protoname]
	return p.CreateClone()
}

// ** MessageBox

type MessageBox struct {
	decochar rune
}

// Use ...
func (mb *MessageBox) Use(s string) {
	length := len(s)
	for i := 0; i < length+4; i++ {
		fmt.Printf("%c", mb.decochar)
	}
	fmt.Println("")
	fmt.Printf("%c %s %c\n", mb.decochar, s, mb.decochar)
	for i := 0; i < length+4; i++ {
		fmt.Printf("%c", mb.decochar)
	}
	fmt.Println("")
}

// CreateClone ...
func (mb *MessageBox) CreateClone() Product {
	return mb.Clone().(*MessageBox)
}

// Clone ...
func (mb *MessageBox) Clone() interface{} {
	copied := *mb
	return &copied
}

// ** UnderlinePen

type UnderlinePen struct {
	ulchar rune
}

// Use ...
func (up *UnderlinePen) Use(s string) {
	length := len(s)
	fmt.Printf("\"%s\"\n", s)
	fmt.Print(" ")
	for i := 0; i < length; i++ {
		fmt.Printf("%c", up.ulchar)
	}
	fmt.Println(" ")
}

// CreateClone ...
func (up *UnderlinePen) CreateClone() Product {
	return up.Clone().(*UnderlinePen)
}

// Clone ...
func (up *UnderlinePen) Clone() interface{} {
	copied := *up
	return &copied
}
