package event

var EventType_Resize string = "resize event"

type ResizeEvent struct {
	Width, Height int32
}

func (k *ResizeEvent) Type() string {
	return EventType_Resize
}

func NewResizeEvent(width, height int32) *ResizeEvent {
	resizeEvent := &ResizeEvent{width, height}

	return resizeEvent
}
