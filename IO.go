package bl

import (
	"xel"
	"bellina/event"
)

func IO_onKey(key xel.Key, action xel.Action) {
	keyEvent := event.NewKeyEvent(key, action)

	event.Fire(keyEvent)
}

func IO_onMouseMove(x,y int32) {
	Mouse_X, Mouse_y = x, y

	node := GetNodeAt(x,y)

	var e *MouseMoveEvent

	if node != nil {
		e = NewMouseMoveEvent(x, y, node)
	}

	for {
		if node == nil {
			break
		}

		if node.OnMouseMoveCallbacks != nil {
			for element := node.OnMouseMoveCallbacks.Front(); element != nil; element = element.Next() {
			    	cb := element.Value.(func(*MouseMoveEvent))

				cb(e)

			}
		}

		if !e.BubbleToParent {
			break
		}

		node = node.Parent

		e.CurrentTarget = node
	}

	event.Fire(e)
}

func IO_onMouseButton(button xel.MouseButton, action xel.Action) {
	node := GetNodeAt(Mouse_X, Mouse_y)

	var e *MouseButtonEvent

	if node != nil {
		e = NewMouseButtonEvent(button, action, node)
	}

	for {
		if node == nil {
			break
		}

		if node.OnMouseButtonCallbacks != nil {
			for element := node.OnMouseButtonCallbacks.Front(); element != nil; element = element.Next() {
				cb := element.Value.(func(*MouseButtonEvent))

				cb(e)
			}
		}

		if !e.BubbleToParent {
			break
		}

		node = node.Parent

		e.CurrentTarget = node
	}

	event.Fire(e)
}

func IO_onResize(width, height int32) {
	resizeEvent := event.NewResizeEvent(width, height)

	event.Fire(resizeEvent)
}


