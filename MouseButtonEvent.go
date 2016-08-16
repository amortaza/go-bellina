package bl

var EventType_Mouse_Button string = "mouse button event"

type MouseButtonEvent struct {
	Button         MouseButton
	ButtonAction   ButtonAction
	Target         *Node // the node the event originated on
	CurrentTarget  *Node
	BubbleToParent bool
}

func (m *MouseButtonEvent) Type() string {
	return EventType_Mouse_Button
}

func NewMouseButtonEvent(button MouseButton, action ButtonAction, target *Node) *MouseButtonEvent {
	e := &MouseButtonEvent{button, action, target, target, true}

	return e
}
