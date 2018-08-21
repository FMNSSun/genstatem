package main

//go:generate genstatem -in=desc.json -out=statemachine.go -package=main


import "fmt"

type InternalState interface {
	Signal(e Event, s State)
	Start(e Event, s State)
	Dump(e Event, s State)
}

type internalState struct {
	counter int
}

func (is *internalState) Signal(e Event, s State) {
	is.counter++
}

func (is *internalState) Start(e Event, s State) {
	is.counter = 0
}

func (is *internalState) Dump(e Event, s State) {
	fmt.Printf("Counter := %d\n", is.counter)
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
