package bl

import "container/list"

var g_longTerm_callbacksByEventType map[string] *list.List
var g_shortTerm_callbacksByEventType map[string] *list.List

func init() {

	g_longTerm_callbacksByEventType = make(map[string] *list.List)
}

type Event interface {

	Type() string
}

func RegisterLongTerm(eventType string, cb func(Event)) {

	callbacks, found := g_longTerm_callbacksByEventType[eventType]

	if !found {

		callbacks = list.New()

		g_longTerm_callbacksByEventType[eventType] = callbacks
	}

	callbacks.PushBack(cb)
}

func RegisterShortTerm(eventType string, callback func(Event)) {

	callbacks, found := g_shortTerm_callbacksByEventType[eventType]

	if !found {

		callbacks = list.New()

		g_shortTerm_callbacksByEventType[eventType] = callbacks
	}

	callbacks.PushBack(callback)
}

func FireEvent(event Event) {

	callbacks, found := g_longTerm_callbacksByEventType[event.Type()]

	if found {

		for e := callbacks.Front(); e != nil; e = e.Next() {

			callback := e.Value.(func(Event))

			callback(event.(Event))
		}
	}

	callbacks2, found2 := g_shortTerm_callbacksByEventType[event.Type()]

	if found2 {

		for e := callbacks2.Front(); e != nil; e = e.Next() {

			callback := e.Value.(func(Event))

			callback(event.(Event))
		}
	}
}
