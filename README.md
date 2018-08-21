# Documentation

This is a tool that can be used to create code from a description of a state machine. The state machine 
is described using a simple JSON structure.

## Format

```
{"states":
  [
    {"name":"idle", "on":"InIdle",
     "transitions": 
       [
         {"event":"start","to":"running","action":"Foo"}
       ]},
    {"name":"running", "on":"InRunning",
     "transitions":
       [
         {"event":"stop","to":"idle","action":"Foo"}
       ]}
  ], 
 "init":"idle",
 "name":"StateMachine",
 "iface":""}
```

### states

An array of states defining the states of the state machine.

#### states.name

Name of the state.

#### states.on

Name of the function to be executed whenever a transition into this state occurse.
The type of the function must be `func(e Event, s State) error`.

#### states.transitions

An array of transitions describing how and when to switch states. 

#### states.transitions.event

The name of the event.

#### states.transitions.to

Name of the target state. If the specified event occurs the state machine will transition
into the target state. This may be left empty if no transition shall occur. 

#### states.transitions.action

Name of the function to be executed whenever this transition occurs. The function is called
before the state is updated. The type of the function must be `func(e Event, s State) error`. This
may be left empty if no function should be executed.

#### states.transitions.condition

Name of a function to be executed to check whether the transition should happen or not. The type
of the function must be `func(e Event, s State) (bool, error)`. If this function returns false
then `action` doesn't happen and no state transition happens. 

### init

The initial state. 

### name

The name of the state machine. This will also be used as the name of the type of the state machine in the
generated code. 

### iface

When `iface` is not empty the state machine will have an internal state of this type. All callbacks
will be invoked on the internal state. 

## Using the generated code

To use the state machine in the generated code use the `New<name you provided>` method. Then you can trigger transitions
be invoking the `Event` method on it:

```go
sm := NewStateMachine()

err := sm.Event(EventStart)

if err != nil {
	panic(err.Error())
}
```

The `Event` function has the type `func(e Event) error`. It may return an error if there's no transition registered 
for this event in the current state. If there's an error in a callback (such as `action`, `condition` or `on`) then `Event` will
abort and return that error. 

The tool will also generate constants `Event*` and `State*` for each event and state. To access the current state of the
state machine use `sm.State()` which has type `func() State`. It also generates a `SetState(state State, event Event, invokeOn bool) error`
method and a `SetIface(iface Iface)` method (if `iface` is used). The special constant `NoEvent` may be used for invalid events. 

For a working example see the `example/` and `example2/` directories.
