package bl

import (
	"xel"
	"bellina/event"
)

func IO_onKey(key xel.Key, action xel.Action) {
	keyEvent := NewKeyEvent(key, action)

	event.Fire(keyEvent)
}

func IO_onMouseMove(x,y int32) {
	Mouse_X, Mouse_y = x, y

	node := GetNodeAt(x,y)

	if node == nil {
		node = Root_Node
	}

	var e *MouseMoveEvent

	if node != nil {
		e = NewMouseMoveEvent(x, y, node)
	}

	for {
		if node == nil {
			break
		}

		node.CallMouseMoveCallbacks(e)

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

	if node == nil {
		node = Root_Node
	}

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


