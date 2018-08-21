package main

//go:generate genstatem -in=desc.json -out=statemachine.go -package=main


import "fmt"

type InternalState interface {
	Signal(e Event, s State) error
	Start(e Event, s State) error
	Dump(e Event, s State) error
	Foobar(e Event, s State) (bool, error)
}

type internalState struct {
	counter int
}

func (is *internalState) Signal(e Event, s State) error {
	is.counter++
	return nil
}

func (is *internalState) Start(e Event, s State) error {
	is.counter = 0
	return nil
}

func (is *internalState) Dump(e Event, s State) error {
	fmt.Printf("Counter := %d\n", is.counter)
	return nil
}

func (is *internalState) Foobar(e Event, s State) (bool, error) {
	return is.counter < 5, nil
}

func main() {
	sm := NewStateMachine(&internalState{})

	for j := 0; j < 3; j++ {
		sm.Event(EventStart)

		for i := 0; i < 10; i++ {
			sm.Event(EventSignal)
		}

		sm.Event(EventStop)
	}
	
}
