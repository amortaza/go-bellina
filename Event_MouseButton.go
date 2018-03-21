package bl

import (
	"container/list"
	"github.com/amortaza/go-hal"
)

var EventType_Mouse_Button = "mouse-button-event"

type MouseButtonEvent struct {

	Button         	hal.MouseButton
	ButtonAction   	hal.ButtonAction

	X, Y			int

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

	e := &MouseButtonEvent{button, action, mouseX, mouseY, target}

	return e
}

func registerOnMouseButtonOnNode(node *Node, cb func(*MouseButtonEvent)) {

	if node.On_Mouse_Button_Callbacks == nil {
		node.On_Mouse_Button_Callbacks = list.New()
	}

	node.On_Mouse_Button_Callbacks.PushBack(cb);
}

