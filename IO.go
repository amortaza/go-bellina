package bl

func io_onKey(key KeyboardKey, action ButtonAction, alt, ctrl, shift bool) {
	keyEvent := NewKeyEvent(key, action, alt, ctrl, shift)

	FireEvent(keyEvent)
}

func io_onMouseMove(x,y int) {
	Mouse_X, Mouse_Y = x, y

	node := getNodeAt__VisibleToMouseEvents(x, y)

	if node == nil {
		node = Root_Node
	}

	e:= NewMouseMoveEvent(x, y, node)

	node.CallMouseMoveCallbacks(e)

	FireEvent(e)
}

func io_onMouseButton(button MouseButton, action ButtonAction) {
	node := getNodeAt__VisibleToMouseEvents(Mouse_X, Mouse_Y)

	if node == nil {
		node = Root_Node
	}

	e := NewMouseButtonEvent(button, Mouse_X, Mouse_Y, action, node)

	if node.OnMouseButtonCallbacks != nil {
		for element := node.OnMouseButtonCallbacks.Front(); element != nil; element = element.Next() {
			cb := element.Value.(func(*MouseButtonEvent))

			cb(e)
		}
	}

	FireEvent(e)
}

func io_onWindowResize(width, height int) {
	resizeEvent := NewWindowResizeEvent(width, height)

	Window_Width, Window_Height = width, height

	FireEvent(resizeEvent)
}


