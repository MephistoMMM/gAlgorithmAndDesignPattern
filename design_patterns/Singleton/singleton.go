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
package singleton

import (
	"sync"
	"sync/atomic"
)

// 1. use Counter to guarantee counter won't be instanced out of package
// use id method to guarantee only struct in this package could implement
// Counter interface
type Counter interface {
	id() *counter

	Add() uint64
	Reset()
	Sub() uint64
	Get() uint64
}

type counter struct {
	count uint64
}

func (c *counter) id() *counter {
	return c
}

// Add return value after count plus 1
func (c *counter) Add() uint64 {
	return atomic.AddUint64(&c.count, 1)
}

// Reset set count to 0
func (c *counter) Reset() {
	atomic.StoreUint64(&c.count, 0)
}

// Sub ...
func (c *counter) Sub() uint64 {
	return atomic.AddUint64(&c.count, (^uint64(1))+1)
}

// Get ...
func (c *counter) Get() uint64 {
	return atomic.LoadUint64(&c.count)
}

var singleCounter *counter

// 2. use sync.Once to make SingleCounter init only once
var once sync.Once

func GetCounter() Counter {
	if singleCounter == nil {
		once.Do(func() {
			singleCounter = &counter{
				count: 0,
			}
		})
	}
	return singleCounter
}
