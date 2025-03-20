package statemachine

import (
	"context"
	"fmt"
)

// State represents the current state of the state machine.
type State string

// StateTransition represents a transition between states in the state machine.
type StateTransition interface {
	GetState() State
	SetState(state State)
}

// Action represents the action to be performed in a state.
type Action func(ctx context.Context, args StateTransition) (StateTransition, error)

// StateMachine represents a generic state machine.
type StateMachine struct {
	states map[State]Action
}

// New creates a new state machine.
func New() *StateMachine {
	return &StateMachine{
		states: make(map[State]Action),
	}
}

// RegisterState registers a state with its corresponding action.
func (sm *StateMachine) RegisterState(state State, action Action) {
	sm.states[state] = action
}

// Run executes the state machine starting from the initial state.
func (sm *StateMachine) Run(ctx context.Context, initialState State, args StateTransition) (StateTransition, error) {
	currentState := initialState
	currentArgs := args

	for {
		action, exists := sm.states[currentState]
		if !exists {
			return currentArgs, fmt.Errorf("state %s not registered", currentState)
		}

		result, err := action(ctx, currentArgs)
		if err != nil {
			return result, err
		}

		nextState := result.GetState()
		if nextState == currentState {
			break
		}

		// Validate that the next state is registered
		if _, exists := sm.states[nextState]; !exists {
			return result, fmt.Errorf("next state %s not registered", nextState)
		}

		currentState = nextState
		currentArgs = result
	}

	return currentArgs, nil
}
