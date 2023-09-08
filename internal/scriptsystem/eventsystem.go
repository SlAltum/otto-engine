package scriptsystem

import (
	"sync"

	"github.com/robertkrimen/otto"

	"example.com/otto-test/internal/logger"
)

var (
	eventLogger *logger.Logger = logger.NewLogger("EVENTSYSTEM")
)

type Listener struct {
	filePath       string
	notifyFuncName string
}

type Event struct {
	name      string
	listeners []*Listener
}

func NewEvent(
	eventName string,
) *Event {
	event := &Event{
		name:      eventName,
		listeners: make([]*Listener, 0),
	}

	return event
}

func (e *Event) addlistener(filePath, notifyFuncName string) {
	listener := &Listener{
		filePath:       filePath,
		notifyFuncName: notifyFuncName,
	}
	e.listeners = append(e.listeners, listener)
}

func (e *Event) Notify(argumentList ...interface{}) {
	for _, listener := range e.listeners {
		vm, ok := vmList.Load(listener.filePath)
		if !ok {
			eventLogger.Log(logger.Error, "load vm fail %s", listener.filePath)
			continue
		}
		_, err := vm.(*otto.Otto).Call(listener.notifyFuncName, nil, argumentList)
		if err != nil {
			eventLogger.Log(logger.Error, "call notify function fail %s", listener.notifyFuncName)
			continue
		}
	}
}

type EventSystem struct {
	events sync.Map
}

func NewEventSystem() *EventSystem {
	eventSystem := &EventSystem{}
	PutApiFuncExport("addlistener", func(call otto.FunctionCall) otto.Value {
		eventName, _ := call.Argument(0).ToString()
		filePath, _ := call.Argument(1).ToString()
		notifyFuncName, _ := call.Argument(2).ToString()
		e, ok := eventSystem.events.Load(eventName)
		if !ok {
			return otto.FalseValue()
		}
		e.(*Event).addlistener(filePath, notifyFuncName)
		return otto.TrueValue()
	})
	return eventSystem
}
