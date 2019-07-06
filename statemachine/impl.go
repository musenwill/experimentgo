package statemachine

import (
	"fmt"
	"sync"
)

var _ StateMachine = &stateContext{}

type stateContext struct {
	sync.Mutex
	currentState State
	exitState    State
	table        map[State]map[Event]Instruct
}

// New function to create a new statemachine
func New(exitState State) StateMachine {
	return &stateContext{
		exitState: exitState,
		table:     map[State]map[Event]Instruct{},
	}
}

// InitState method to set the first state of the statemachine
func (s *stateContext) InitState(state State) error {
	return s.setState(state)
}

// CurrentState method to get the current state ot the statemachine
func (s *stateContext) CurrentState() State {
	s.Lock()
	defer s.Unlock()

	return s.currentState
}

// Register a handler to the statemachine, tell it what to do when an event is sent to the
// statemachine on a given state
func (s *stateContext) RegisterInstruct(state State, event Event, instruct Instruct) error {
	s.Lock()
	defer s.Unlock()

	eventTable, ok := s.table[state]
	if !ok {
		eventTable = map[Event]Instruct{}
		s.table[state] = eventTable
	}

	_, ok = eventTable[event]
	if ok {
		// log warning, new instruct will overwrite the exist one
	}

	eventTable[event] = instruct

	return nil
}

// Trigger used to send a event and param data to the statemachine
func (s *stateContext) Trigger(event Event, data interface{}) error {
	currentState := s.CurrentState()

	// if no event registered for this state, then this may be the end state and just do nothing
	eventTable, ok := s.table[currentState]
	if !ok {
		return nil
	}

	instruct, ok := eventTable[event]

	// check condition
	if instruct.Condition != nil {
		if err := instruct.Condition(currentState, event, data); err != nil {
			return err
		}
	}

	// do action
	if instruct.Action != nil {
		if err := instruct.Action(currentState, event, data); err != nil {
			return err
		}
	}

	// change to new state
	return s.setState(instruct.NewState)
}

func (s *stateContext) setState(state State) error {
	if state.IsNone() {
		return nil
	}

	s.Lock()
	defer s.Unlock()

	if s.exitState == state {
		s.currentState = state
		return nil
	}

	if _, ok := s.table[state]; !ok {
		return fmt.Errorf("unsupported state %s", state)
	}
	s.currentState = state

	return nil
}
