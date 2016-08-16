package bl

var EventType_Resize string = "resize event"

type ResizeEvent struct {
	Width, Height int
}

func (k *ResizeEvent) Type() string {
	return EventType_Resize
}

func NewResizeEvent(width, height int) *ResizeEvent {
	resizeEvent := &ResizeEvent{width, height}

	return resizeEvent
}
