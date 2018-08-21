package main

//go:generate genstatem -in=desc.json -out=statemachine.go -package=main


import "fmt"

func InIdle(event Event, state State) error {
	fmt.Println("InIdle")
	return nil
}

func InRunning(event Event, state State) error {
	fmt.Println("InRunning")
	return nil
}

func Foo(event Event, state State) error {
	fmt.Println("Foo")
	return nil
}

func main() {
	sm := NewStateMachine()

	err := sm.Event(EventStart)

	if err != nil {
		panic(err.Error())
	}

	err = sm.Event(EventStop)

	if err != nil {
		panic(err.Error())
	}
}
