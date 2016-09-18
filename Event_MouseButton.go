package bl

import "container/list"

var EventType_Mouse_Button string = "mouse-button-event"

type MouseButtonEvent struct {
	Button         	MouseButton
	X, Y		int
	ButtonAction   	ButtonAction
	Target         	*Node
}

func (m *MouseButtonEvent) Type() string {
	return EventType_Mouse_Button
}

const (
	Mouse_Button_Left MouseButton = 1 + iota
	Mouse_Button_Right
)

const (
	Button_Action_Down ButtonAction = 1 + iota
	Button_Action_Up
)

func NewMouseButtonEvent(button MouseButton, mouseX, mouseY int, action ButtonAction, target *Node) *MouseButtonEvent {
	e := &MouseButtonEvent{button, mouseX, mouseY, action, target}

	return e
}

func onMouseButtonOnNode(node *Node, cb func(*MouseButtonEvent)) {
	if node.OnMouseButtonCallbacks == nil {
		node.OnMouseButtonCallbacks = list.New()
	}

	node.OnMouseButtonCallbacks.PushBack(cb);
}

