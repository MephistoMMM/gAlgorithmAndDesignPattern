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
package state

// Context interface of state design pattern
type Context interface {
	SetClock(hour int)
	ChangeState(state State)
	CallSecurityCenter(msg string)
	RecordLog(msg string)
}

// State interface of state design pattern
type State interface {
	DoClock(ctx Context, hour int)
	DoUse(ctx Context)
	DoAlarm(ctx Context)
	DoPhone(ctx Context)
	String() string
}

var (
	// NightState is the state represents night
	NightState State
	// DayState is the state represents day
	DayState State
)

func init() {
	NightState = nightState{}
	DayState = dayState{}
}

type dayState struct{}

// DoClock ...
func (ds dayState) DoClock(ctx Context, hour int) {
	if hour < 9 || 17 <= hour {
		ctx.ChangeState(NightState)
	}
}

// DoUse ...
func (ds dayState) DoUse(ctx Context) {
	ctx.RecordLog("Use GoldRepo (Day)")
}

// DoPhone ...
func (ds dayState) DoAlarm(ctx Context) {
	ctx.CallSecurityCenter("Press Alarm Bell (Day)")
}

// DoPhone ...
func (ds dayState) DoPhone(ctx Context) {
	ctx.CallSecurityCenter("Call (Day)")
}

// String ...
func (ds dayState) String() string {
	return "[DAY]"
}

type nightState struct{}

// DoClock ...
func (ds nightState) DoClock(ctx Context, hour int) {
	if hour >= 9 && 17 > hour {
		ctx.ChangeState(DayState)
	}
}

// DoUse ...
func (ds nightState) DoUse(ctx Context) {
	ctx.CallSecurityCenter("Emergency! Use GoldRope at Night")
}

// DoPhone ...
func (ds nightState) DoAlarm(ctx Context) {
	ctx.CallSecurityCenter("Press Alarm Bell (Night)")
}

// DoPhone ...
func (ds nightState) DoPhone(ctx Context) {
	ctx.CallSecurityCenter("Record (Night)")
}

// String ...
func (ds nightState) String() string {
	return "[NIGHT]"
}
