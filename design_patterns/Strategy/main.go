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
	"gAaD/design_patterns/Strategy/player"
	"gAaD/design_patterns/Strategy/strategy"
	"time"
)

func main() {
	seed1, seed2 := int64(time.Now().Unix()), int64(time.Now().UnixNano()>>10)
	player1 := player.NewPlayer("Taro", strategy.NewWinningStrategy(seed1))
	player2 := player.NewPlayer("Hana", strategy.NewProbStrategy(seed2))
	for i := 0; i < 10000; i++ {
		nextHand1 := player1.NextHand()
		nextHand2 := player2.NextHand()
		if nextHand1.IsStrongerThan(nextHand2) {
			player1.Win()
			player2.Lose()
		} else if nextHand2.IsStrongerThan(nextHand1) {
			player2.Win()
			player1.Lose()
		} else {
			player1.Even()
			player2.Even()
		}
	}

	fmt.Println(player1.String())
	fmt.Println(player2.String())
}
