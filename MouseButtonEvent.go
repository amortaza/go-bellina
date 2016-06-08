package bl

import "xel"

var Mouse_Button_Event_Type string = "mouse button event"

type MouseButtonEvent struct {
	Button xel.MouseButton
	Action xel.Action
	Target *Node // the node the event originated on
	CurrentTarget *Node
	BubbleToParent bool
}

func (m *MouseButtonEvent) Type() string {
	return Mouse_Button_Event_Type
}

func NewMouseButtonEvent(button xel.MouseButton, action xel.Action, target *Node) *MouseButtonEvent {
	e := &MouseButtonEvent{button, action, target, target, true}

	return e
}
