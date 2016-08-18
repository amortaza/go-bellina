package bl

type ButtonAction int
type MouseButton int
type KeyboardKey int

func IO_onKey(key KeyboardKey, action ButtonAction, alt, ctrl, shift bool) {
	keyEvent := NewKeyEvent(key, action, alt, ctrl, shift)

	FireEvent(keyEvent)
}

func IO_onMouseMove(x,y int) {
	Mouse_X, Mouse_Y = x, y

	node := GetNodeAt__VisibleToEvents(x, y)

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

		if !e.BubbleToParent || node.PreventBubbling {
			break
		}

		node = node.Parent

		e.CurrentTarget = node
	}

	FireEvent(e)
}

func IO_onMouseButton(button MouseButton, action ButtonAction) {
	node := GetNodeAt__VisibleToEvents(Mouse_X, Mouse_Y)

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

		if !e.BubbleToParent || node.PreventBubbling {
			break
		}

		node = node.Parent

		e.CurrentTarget = node
	}

	FireEvent(e)
}

func IO_onResize(width, height int) {
	resizeEvent := NewResizeEvent(width, height)

	FireEvent(resizeEvent)
}


