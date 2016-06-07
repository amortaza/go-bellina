package core

import (
	"xel"
	"bellina/event"
)

func IO_onKey(key xel.Key, action xel.Action) {
	keyEvent := event.NewKeyEvent(key, action)

	event.Fire(keyEvent)
}

func IO_onMouseMove(x,y int) {
	mouseMoveEvent := event.NewMouseMoveEvent(x, y)

	event.Fire(mouseMoveEvent)
}

func IO_onMouseButton(button xel.MouseButton, action xel.Action) {
	mouseButtonEvent := event.NewMouseButtonEvent(button, action)

	event.Fire(mouseButtonEvent)
}

func IO_onResize(width,height int32) {
	resizeEvent := event.NewResizeEvent(width, height)

	event.Fire(resizeEvent)
}


