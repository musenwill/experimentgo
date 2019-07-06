package statemachine

// State is the name of states
type State string

// Event is sent to statemachine, which will response to it and transfer to another state
type Event string

// StateMachine is an interface defines the behaves of a statemachine
type StateMachine interface {
	// Tell what's the first state of the statemachine
	InitState(state State) error
	// Get current state of the statemachine
	CurrentState() State
	// Register a handler to the statemachine, tell it what to do when an event is sent to the
	// statemachine on a given state
	RegisterInstruct(state State, event Event, instruct Instruct) error
	// Trigger used to send a event and param data to the statemachine
	Trigger(event Event, data interface{}) error
}

// Instruct is an interface defines the reaction of an event to a statemchine
type Instruct struct {
	Condition func(state State, event Event, data interface{}) error
	Action    func(state State, event Event, data interface{}) error
	NewState  State
}

// IsNone check if State is a none value
func (s State) IsNone() bool {
	return len(s) <= 0
}
