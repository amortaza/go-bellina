package bl

import (
	"container/list"
	"github.com/amortaza/go-hal"
)

var EventType_Mouse_Button string = "mouse-button-event"

type MouseButtonEvent struct {
	Button         	hal.MouseButton
	X, Y			int
	ButtonAction   	hal.ButtonAction
	Target         	*Node
}

func (m *MouseButtonEvent) Type() string {
	return EventType_Mouse_Button
}

func NewMouseButtonEvent(
		button hal.MouseButton,
		mouseX, mouseY int,
		action hal.ButtonAction,
		target *Node) *MouseButtonEvent {

	e := &MouseButtonEvent{button, mouseX, mouseY, action, target}

	return e
}

func onMouseButtonOnNode(node *Node, cb func(*MouseButtonEvent)) {

	if node.OnMouseButtonCallbacks == nil {
		node.OnMouseButtonCallbacks = list.New()
	}

	node.OnMouseButtonCallbacks.PushBack(cb);
}

