package bl

import "container/list"

var EventType_Mouse_Button string = "mouse-button-event"

type MouseButtonEvent struct {
	Button         MouseButton
	ButtonAction   ButtonAction
	Target         *Node
}

func (m *MouseButtonEvent) Type() string {
	return EventType_Mouse_Button
}

func NewMouseButtonEvent(button MouseButton, action ButtonAction, target *Node) *MouseButtonEvent {
	e := &MouseButtonEvent{button, action, target}

	return e
}

func onMouseButtonOnNode(node *Node, cb func(*MouseButtonEvent)) {
	if node.OnMouseButtonCallbacks == nil {
		node.OnMouseButtonCallbacks = list.New()
	}

	node.OnMouseButtonCallbacks.PushBack(cb);
}

