package bl

var Mouse_Move_Event_Type string = "mouse move event"

type MouseMoveEvent struct {
	X, Y int32
	Target *Node // the node the event originated on
	CurrentTarget *Node
	BubbleToParent bool
}

func (m *MouseMoveEvent) Type() string {
	return Mouse_Move_Event_Type
}

func NewMouseMoveEvent(x, y int32, target *Node) *MouseMoveEvent {
	e := &MouseMoveEvent{x, y, target, target, true}

	return e
}
