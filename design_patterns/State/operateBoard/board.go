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
package operateBoard

import (
	"fmt"
	"github.com/MephistoMMM/gAlgorithmAndDesignPattern/design_patterns/State/state"
	"time"
)

// Order is order passed to operateChan
type Order uint8

const (
	// DO_USE use order
	DO_USE Order = iota
	// DO_ALARM alarm order
	DO_ALARM
	// DO_PHONE phone order
	DO_PHONE
	// do_STOP stop OperateBoard
	do_STOP
)

// status is the inner status of OperateBoard
type status uint8

const (
	sRUNNING status = iota
	sSTOPPED
)

// OperateBoard implement Context interface
type OperateBoard struct {
	state       state.State
	operateChan chan Order

	status status
}

func NewOperateBoard() *OperateBoard {
	return &OperateBoard{
		state:       state.DayState,
		operateChan: make(chan Order),
		status:      sSTOPPED,
	}
}

// SetClock ...
func (ob *OperateBoard) SetClock(hour int) {
	ob.state.DoClock(ob, hour)
}

// ChangeState ...
func (ob *OperateBoard) ChangeState(state state.State) {
	fmt.Printf("Change state from %s to %s \n", ob.state, state)
	ob.state = state
}

// CallSecurityCenter ...
func (ob *OperateBoard) CallSecurityCenter(msg string) {
	fmt.Println("Call! " + msg)
}

// RecordLog ...
func (ob *OperateBoard) RecordLog(msg string) {
	fmt.Println("Record... " + msg)
}

// Start ...
func (ob *OperateBoard) Start() error {
	if ob.status == sRUNNING {
		return nil
	}

	go func() {
		fmt.Println("OperateBoard is started.")
		ob.status = sRUNNING
		c := time.Tick(1 * time.Second)
	FINISH:
		for {
			select {
			case t := <-c:
				ob.SetClock(t.Second() % 24)
			case order := <-ob.operateChan:
				if order == do_STOP {
					ob.status = sSTOPPED
					break FINISH
				}
				ob.operate(order)
			}
		}
		fmt.Println("OperateBoard is stopped.")
	}()

	return nil
}

// Stop ...
func (ob *OperateBoard) Stop() error {
	ob.operateChan <- do_STOP
	return nil
}

// operate ...
func (ob *OperateBoard) operate(op Order) {
	switch op {
	case DO_USE:
		ob.state.DoUse(ob)
	case DO_ALARM:
		ob.state.DoAlarm(ob)
	case DO_PHONE:
		ob.state.DoPhone(ob)
	}
}

// OperateChan ...
func (ob *OperateBoard) OperateChan() chan<- Order {
	return ob.operateChan
}
