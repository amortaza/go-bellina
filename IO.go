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

	e:= NewMouseMoveEvent(x, y, node)

	node.CallMouseMoveCallbacks(e)

	FireEvent(e)
}

func IO_onMouseButton(button MouseButton, action ButtonAction) {
	node := GetNodeAt__VisibleToEvents(Mouse_X, Mouse_Y)

	if node == nil {
		node = Root_Node
	}

	e := NewMouseButtonEvent(button, action, node)

	if node.OnMouseButtonCallbacks != nil {
		for element := node.OnMouseButtonCallbacks.Front(); element != nil; element = element.Next() {
			cb := element.Value.(func(*MouseButtonEvent))

			cb(e)
		}
	}

	FireEvent(e)
}

func IO_onResize(width, height int) {
	resizeEvent := NewResizeEvent(width, height)

	FireEvent(resizeEvent)
}


