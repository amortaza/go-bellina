package event

import "container/list"

var g_registerLongTermCallbacksByEventType map[string] *list.List
var G_registerShortTermCallbacksByEventType map[string] *list.List

func init() {
	g_registerLongTermCallbacksByEventType = make(map[string] *list.List)
}

type Event interface {
	Type() string
}

func RegisterLongTerm(eventType string, callback func(Event)) {
	callbacks, found := g_registerLongTermCallbacksByEventType[eventType]

	if !found {
		callbacks = list.New()

		g_registerLongTermCallbacksByEventType[eventType] = callbacks
	}

	callbacks.PushBack(callback)
}

func RegisterShortTerm(eventType string, callback func(Event)) {
	callbacks, found := G_registerShortTermCallbacksByEventType[eventType]

	if !found {
		callbacks = list.New()

		G_registerShortTermCallbacksByEventType[eventType] = callbacks
	}

	callbacks.PushBack(callback)
}

func Fire(event Event) {
	callbacks, found := g_registerLongTermCallbacksByEventType[event.Type()]

	if found {
		for e := callbacks.Front(); e != nil; e = e.Next() {
			callback := e.Value.(func(Event))

			callback(event.(Event))
		}
	}

	callbacks2, found2 := G_registerShortTermCallbacksByEventType[event.Type()]

	if found2 {
		for e := callbacks2.Front(); e != nil; e = e.Next() {
			callback := e.Value.(func(Event))

			callback(event.(Event))
		}
	}
}
