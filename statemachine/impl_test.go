package statemachine

import (
	"fmt"
	"testing"
)

/*
This example monitors process status on an operating system.
Any process has four states: ready, runing, suspend and exit.
1. On ready state, if receive 'cpu free' event, then process go into runing state.
2. On runing state, if receive 'cpu busy' event, then process go into suspend state.
3. On suspend state, if receive 'priority high' event, then process go into ready state.
4. On any state, if receive 'kill' event, then process go into exit state.
5. exit state receives no event.
*/

var state = struct {
	READY, RUNING, SUSPEND, EXIT State
}{
	"READY", "RUNING", "SUSPEND", "EXIT",
}

var event = struct {
	CPUFree, CPUBusy, PriorityHigh, Kill Event
}{
	"CPU_FREE", "CPU_BUSY", "PRIORITY_HIGH", "KILL",
}

func new() StateMachine {
	stateMachine := New(state.EXIT)
	stateMachine.RegisterInstruct(state.READY, event.CPUFree, Instruct{Condition: nil, Action: commonAction, NewState: state.RUNING})
	stateMachine.RegisterInstruct(state.READY, event.Kill, Instruct{Condition: nil, Action: commonAction, NewState: state.EXIT})

	stateMachine.RegisterInstruct(state.RUNING, event.CPUBusy, Instruct{Condition: nil, Action: commonAction, NewState: state.SUSPEND})
	stateMachine.RegisterInstruct(state.READY, event.Kill, Instruct{Condition: nil, Action: commonAction, NewState: state.EXIT})

	stateMachine.RegisterInstruct(state.SUSPEND, event.PriorityHigh, Instruct{Condition: nil, Action: commonAction, NewState: state.READY})
	stateMachine.RegisterInstruct(state.SUSPEND, event.Kill, Instruct{Condition: nil, Action: commonAction, NewState: state.EXIT})

	stateMachine.RegisterInstruct(state.EXIT, event.Kill, Instruct{Condition: nil, Action: commonAction, NewState: state.EXIT})

	stateMachine.InitState(state.READY)

	return stateMachine
}

func commonAction(state State, event Event, data interface{}) error {
	fmt.Printf("on state %s receive event %s\n", state, event)
	return nil
}

func TestCaseProcessScheduler(t *testing.T) {
	eventList := []Event{event.PriorityHigh, event.CPUBusy, event.CPUFree, event.CPUBusy,
		event.CPUFree, event.PriorityHigh, event.Kill, event.CPUFree, event.CPUBusy}
	expStateList := []State{state.READY, state.READY, state.READY, state.RUNING, state.SUSPEND, state.SUSPEND, state.READY, state.EXIT}
	stateMachine := new()

	index := 0
	for _, e := range eventList {
		if exp, act := expStateList[index], stateMachine.CurrentState(); exp != act {
			t.Errorf("statemachine error, current state: %v != %v", act, exp)
		}
		index++

		if stateMachine.CurrentState() != state.EXIT {
			if err := stateMachine.Trigger(e, nil); err != nil {
				t.Error(err)
			}
		} else {
			break
		}
	}
}
