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
package factory

import "fmt"

// * interface package

type Product interface {
	Use()
}

type Factory interface {
	CreateProduct(owner string) Product
	RegisterProduct(product Product)
}

// template function
func Create(factory Factory, owner string) Product {
	p := factory.CreateProduct(owner)
	factory.RegisterProduct(p)
	return p
}

// * concrete Package

type IDCard struct {
	owner string
}

func NewIDCard(owner string) IDCard {
	fmt.Println("create ID card for " + owner)
	return IDCard{
		owner: owner,
	}
}

// Use ...
func (ic IDCard) Use() {
	fmt.Println("Use " + ic.owner + "'s ID card")
}

// getOwner ...
func (ic IDCard) getOwner() string {
	return ic.owner
}

type IDCardFactory struct {
	owners []string
}

// createProduct ...
func (icf *IDCardFactory) CreateProduct(owner string) Product {
	return NewIDCard(owner)
}

// registerProduct ...
func (icf *IDCardFactory) RegisterProduct(product Product) {
	icf.owners = append(icf.owners, product.(IDCard).getOwner())
}
