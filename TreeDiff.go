package bl

import (
	"bellina/event"
)

type NodeDeleteEvent struct {
	node *Node
}

type NodeCreateEvent struct {
	node *Node
}

var NodeDelete_Event_Type string = "node delete event"
var NodeCreate_Event_Type string = "node create event"

func (k *NodeDeleteEvent) Type() string {
	return NodeDelete_Event_Type
}

func (k *NodeCreateEvent) Type() string {
	return NodeCreate_Event_Type
}

func NewNodeDeleteEvent(node *Node) *NodeDeleteEvent {
	return &NodeDeleteEvent{node}
}

func NewNodeCreateEvent(node *Node) *NodeCreateEvent {
	return &NodeCreateEvent{node}
}

func detectDifferences(prev, now map[string] *Node) {
	for id, node := range prev {
		_, ok := now[id]

		if !ok {
			// node was deleted
			deleteEvent := NewNodeDeleteEvent(node)

			event.Fire(deleteEvent)
		}
	}

	for id, node := range now {
		_, ok := prev[id]

		if !ok {
			// node was created
			createEvent := NewNodeCreateEvent(node)

			event.Fire(createEvent)
		}
	}
}
