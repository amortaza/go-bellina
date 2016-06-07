package event

import "xel"

var Mouse_Button_Event_Type string = "mouse button event"

type MouseButtonEvent struct {
	Button xel.MouseButton
	Action xel.Action
}

func (m *MouseButtonEvent) Type() string {
	return Mouse_Button_Event_Type
}

func NewMouseButtonEvent(button xel.MouseButton, action xel.Action) *MouseButtonEvent {
	mouseButtonEvent := &MouseButtonEvent{button, action}

	return mouseButtonEvent
}
