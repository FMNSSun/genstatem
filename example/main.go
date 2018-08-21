package main

//go:generate genstatem -in=desc.json -out=statemachine.go -package=main


import "fmt"

func InIdle(event Event, state State) {
	fmt.Println("InIdle")
}

func InRunning(event Event, state State) {
	fmt.Println("InRunning")
}

func Foo(event Event, state State) {
	fmt.Println("Foo")
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
