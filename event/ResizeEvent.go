package event

var Resize_Event_Type string = "resize event"

type ResizeEvent struct {
	Width, Height int32
}

func (k *ResizeEvent) Type() string {
	return Resize_Event_Type
}

func NewResizeEvent(width, height int32) *ResizeEvent {
	resizeEvent := &ResizeEvent{width, height}

	return resizeEvent
}
