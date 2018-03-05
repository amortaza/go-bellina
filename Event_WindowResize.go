package bl

var EventType_Window_Resize = "window-resize-event"

type WindowResizeEvent struct {

	Width, Height int
}

func (k *WindowResizeEvent) Type() string {

	return EventType_Window_Resize
}

func NewWindowResizeEvent(width, height int) *WindowResizeEvent {

	event := &WindowResizeEvent{width, height}

	return event
}
