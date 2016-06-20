package bl

var EventType_Mouse_Move string = "mouse move event"

type MouseMoveEvent struct {
	X, Y int32
	Target *Node // the node the event originated on
	CurrentTarget *Node
	BubbleToParent bool
}

func (m *MouseMoveEvent) Type() string {
	return EventType_Mouse_Move
}

func NewMouseMoveEvent(x, y int32, target *Node) *MouseMoveEvent {
	e := &MouseMoveEvent{x, y, target, target, true}

	return e
}
