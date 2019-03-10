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
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Observer is interface of observer in Observer design pattern
type Observer interface {
	Update(ng NumberGenerator)
}

// NumberGenerator is interface of subject in Observer design pattern
type NumberGenerator interface {
	AddObserver(obs Observer)
	DeleteObserver(obs Observer)
	NotifyObservers()
	GetNumber() int
	Execute()
}

// BaseGenerator has default implement methods of NumberGenerator
type BaseGenerator struct {
	observers []Observer
}

func NewBaseGenerator() *BaseGenerator {
	return &BaseGenerator{
		observers: make([]Observer, 0),
	}
}

// AddObserver ...
func (bg *BaseGenerator) AddObserver(obs Observer) {
	bg.observers = append(bg.observers, obs)
}

// DeleteObserver ...
func (bg *BaseGenerator) DeleteObserver(obs Observer) {
	for i, v := range bg.observers {
		if v == obs {
			bg.observers = append(bg.observers[:i], bg.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers ...
func (bg *BaseGenerator) NotifyObservers(ng NumberGenerator) {
	for _, v := range bg.observers {
		v.Update(ng)
	}
}

// RandomNumberGenerator is concrete subject
type RandomNumberGenerator struct {
	*BaseGenerator

	number int
	random *rand.Rand
}

func NewRandomNumberGenerator() NumberGenerator {
	return &RandomNumberGenerator{
		BaseGenerator: NewBaseGenerator(),
		random:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// NotifyObservers ...
func (rng *RandomNumberGenerator) NotifyObservers() {
	rng.BaseGenerator.NotifyObservers(rng)
}

// GetNumber ...
func (rng *RandomNumberGenerator) GetNumber() int {
	return rng.number
}

// Execute ...
func (rng *RandomNumberGenerator) Execute() {
	for i := 0; i < 20; i++ {
		rng.number = rng.random.Intn(50)
		rng.NotifyObservers()
	}
}

type DigitObserver struct {
}

// Update ...
func (do DigitObserver) Update(ng NumberGenerator) {
	fmt.Printf("Digital Observer: %d.\n", ng.GetNumber())
}

type GraphObserver struct {
}

// Update ...
func (gpho GraphObserver) Update(ng NumberGenerator) {
	fmt.Print("Graphal Observer:")
	for i := 0; i < ng.GetNumber(); i++ {
		fmt.Print("*")
	}
	fmt.Println("")
}
