package bl

import "container/list"

var EventType_Mouse_Move = "mouse-move-event"

type MouseMoveEvent struct {

	X, Y int

	Target *Node
}

func (m *MouseMoveEvent) Type() string {

	return EventType_Mouse_Move
}

func NewMouseMoveEvent(x, y int, target *Node) *MouseMoveEvent {

	e := &MouseMoveEvent{x, y, target}

	return e
}

func registerOnMouseMoveOnNode(node *Node, cb func(*MouseMoveEvent)) {

	if node.On_Mouse_Move_Callbacks == nil {

		node.On_Mouse_Move_Callbacks = list.New()
	}

	node.On_Mouse_Move_Callbacks.PushBack(cb);
}

