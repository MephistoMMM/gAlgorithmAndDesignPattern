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
package player

import (
	"fmt"
	"gAaD/design_patterns/Strategy/hand"
	"gAaD/design_patterns/Strategy/strategy"
)

type Player struct {
	name      string
	strategy  strategy.Strategy
	wincount  int
	losecount int
	gamecount int
}

func NewPlayer(name string, strategy strategy.Strategy) *Player {
	return &Player{
		name:     name,
		strategy: strategy,
	}
}

// NextHand  ...
func (pl *Player) NextHand() hand.Hand {
	return pl.strategy.NextHand()
}

// Win  ...
func (pl *Player) Win() {
	pl.strategy.Study(true)
	pl.wincount++
	pl.gamecount++
}

// Lose  ...
func (pl *Player) Lose() {
	pl.strategy.Study(false)
	pl.losecount++
	pl.gamecount++
}

// Even  ...
func (pl *Player) Even() {
	pl.gamecount++
}

// String  ...
func (pl *Player) String() string {
	return fmt.Sprintf("[%s:%d games, %d win, %d lose]", pl.name, pl.wincount, pl.losecount)
}
