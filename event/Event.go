package event

import "container/list"

var g_callbacksByEventType map[string] *list.List

func init() {
	g_callbacksByEventType = make(map[string] *list.List)
}

type Event interface {
	Type() string
}

func Register(eventType string, callback func(Event)) {
	callbacks, found := g_callbacksByEventType[eventType]

	if !found {
		callbacks = list.New()

		g_callbacksByEventType[eventType] = callbacks
	}

	callbacks.PushBack(callback)
}

func Fire(event Event) {
	callbacks, found := g_callbacksByEventType[event.Type()]

	if !found {
		return
	}

	for e := callbacks.Front(); e != nil; e = e.Next() {
		callback := e.Value.(func(Event))

		callback(event.(Event))
	}
}
