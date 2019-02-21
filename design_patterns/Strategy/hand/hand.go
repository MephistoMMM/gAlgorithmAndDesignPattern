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
package hand

// HandValue type of rock , scissors and paper.
type HandValue int

const (
	// ROCK represent rock.
	ROCK HandValue = 0 + iota
	// SCISSORS represent scissors.
	SCISSORS
	// PAPER represent paper.
	PAPER
)

type Hand interface {
	// IsStrongerThan if this win h , return true
	IsStrongerThan(h Hand) bool
	// IsStrongerThan if h win this , return true
	IsWeakerThan(h Hand) bool
	// String ...
	String() string

	value() HandValue
}

var handArr = []Hand{
	&hand{ROCK},
	&hand{SCISSORS},
	&hand{PAPER},
}

var name = []string{
	"rock",
	"scissors",
	"paper",
}

type hand struct {
	handvalue HandValue
}

// value ...
func (hd *hand) value() HandValue {
	return hd.handvalue
}

// fight  count: even 0 , win 1 , lose -1
func (hd *hand) fight(h Hand) int {
	vhd := hd.value()
	vh := h.value()
	if vhd == vh {
		return 0
	}
	if (vhd+1)%3 == vh {
		return 1
	}
	return -1
}

// IsStrongerThan ...
func (hd *hand) IsStrongerThan(h Hand) bool {
	return hd.fight(h) == 1
}

// IsWeakerThan ...
func (hd *hand) IsWeakerThan(h Hand) bool {
	return hd.fight(h) == -1
}

// String  ...
func (hd *hand) String() string {
	return name[int(hd.handvalue)]
}

// Of return Hand represented `handvalue`
func Of(handvalue HandValue) Hand {
	return handArr[int(handvalue)]
}
