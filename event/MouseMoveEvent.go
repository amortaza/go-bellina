package event

var Mouse_Move_Event_Type string = "mouse move event"

type MouseMoveEvent struct {
	x, y int
}

func (m *MouseMoveEvent) Type() string {
	return Mouse_Move_Event_Type
}

func NewMouseMoveEvent(x, y int) *MouseMoveEvent {
	mouseMoveEvent := &MouseMoveEvent{x, y}

	return mouseMoveEvent
}
