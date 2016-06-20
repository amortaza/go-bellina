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

var EventType_NodeDelete string = "node delete event"
var EventType_NodeCreate string = "node create event"

func (k *NodeDeleteEvent) Type() string {
	return EventType_NodeDelete
}

func (k *NodeCreateEvent) Type() string {
	return EventType_NodeCreate
}

func NewNodeRemoveEvent(node *Node) *NodeDeleteEvent {
	return &NodeDeleteEvent{node}
}

func NewNodeAddEvent(node *Node) *NodeCreateEvent {
	return &NodeCreateEvent{node}
}

func detectDifferences(prev, now map[string] *Node) {
	for id, node := range prev {
		_, ok := now[id]

		if !ok {
			// node was deleted
			removeEvent := NewNodeRemoveEvent(node)

			for _, plugin := range g_pluginByName {
				plugin.OnNodeRemoved(node)
			}

			event.Fire(removeEvent)
		}
	}

	for id, node := range now {
		_, ok := prev[id]

		if !ok {
			// node was created
			addEvent := NewNodeAddEvent(node)

			for _, plugin := range g_pluginByName {
				plugin.OnNodeAdded(node)
			}

			event.Fire(addEvent)
		}
	}
}
