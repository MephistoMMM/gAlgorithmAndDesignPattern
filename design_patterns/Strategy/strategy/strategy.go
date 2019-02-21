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
package strategy

import (
	"gAaD/design_patterns/Strategy/hand"
	"math/rand"
)

// Strategy interface of morra game strategy
type Strategy interface {
	// NextHand give next hand.
	NextHand() hand.Hand
	// Study learn previous game result to decide next hand to play.
	Study(win bool)
}

// WinningStrategy is a morra game strategy, it decides to give previous
// hand when winning the previous game
type WinningStrategy struct {
	random   *rand.Rand
	won      bool
	prevHand hand.Hand
}

// NewWinningStrategy create and init winning strategy
func NewWinningStrategy(seed int64) Strategy {
	return &WinningStrategy{
		random: rand.New(rand.NewSource(seed)),
		won:    false,
	}
}

// NextHand  ...
func (ws *WinningStrategy) NextHand() hand.Hand {
	if !ws.won {
		ws.prevHand = hand.Of(hand.HandValue(ws.random.Intn(3)))
	}
	return ws.prevHand
}

// Study  ...
func (ws *WinningStrategy) Study(win bool) {
	ws.won = win
}

// ProbStrategy is a morra game strategy, it decides to give hand
// accroding to winning probability relatived with previous hand
// and current hand.
type ProbStrategy struct {
	random           *rand.Rand
	won              bool
	prevHandValue    int
	currentHandValue int
	history          [3][3]int
}

// NewProbStrategy create and init ProbStrategy
func NewProbStrategy(seed int64) Strategy {
	return &ProbStrategy{
		random: rand.New(rand.NewSource(seed)),
		history: [3][3]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
	}
}

// NextHand  ...
func (ps *ProbStrategy) NextHand() hand.Hand {
	if ps.won {
		return hand.Of(hand.HandValue(ps.currentHandValue))
	}

	current := ps.currentHandValue
	bet := ps.random.Intn(ps.sum(current))
	handvalue := 0
	if bet < ps.history[current][0] {
		handvalue = 0
	} else if bet < ps.history[current][0]+ps.history[current][0] {
		handvalue = 1
	} else {
		handvalue = 2
	}

	ps.prevHandValue = ps.currentHandValue
	ps.currentHandValue = handvalue
	return hand.Of(hand.HandValue(ps.currentHandValue))
}

// sum  ...
func (ps *ProbStrategy) sum(hv int) int {
	sum := 1
	for i := 0; i < 3; i++ {
		sum += ps.history[hv][i]
	}
	return sum
}

// Study  ...
func (ps *ProbStrategy) Study(win bool) {
	if win {
		ps.history[ps.prevHandValue][ps.currentHandValue]++
	} else {

		ps.history[ps.prevHandValue][(ps.currentHandValue+1)%3]++
		ps.history[ps.prevHandValue][(ps.currentHandValue+2)%3]++
	}
	ps.won = win
}
