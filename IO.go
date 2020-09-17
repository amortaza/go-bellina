package bl

import "github.com/amortaza/go-hal"

func io_onKey(
		key hal.KeyboardKey,
		action hal.ButtonAction,
		alt, ctrl, shift bool) {

	keyEvent := NewKeyEvent(key, action, alt, ctrl, shift)

	FireEvent(keyEvent)
}

func io_onMouseMove(x, y int) {
	if Root_Node == nil {
		return
	}

	Mouse_X, Mouse_Y = x, y

	node := g_spacial.GetNodeAt_VisibleToMouseEvents(x, y)

	if node == nil {
		node = Root_Node
	}

	e:= NewMouseMoveEvent(x, y, node)

	node.CallMouseMoveCallbacks(e)

	FireEvent(e)
}

func io_onMouseButton(
		button hal.MouseButton,
		action hal.ButtonAction) {

	node := g_spacial.GetNodeAt_VisibleToMouseEvents(Mouse_X, Mouse_Y)

	if node == nil {
		node = Root_Node
	}

	e := NewMouseButtonEvent(button, Mouse_X, Mouse_Y, action, node)

	if node.On_Mouse_Button_Callbacks != nil {
		for element := node.On_Mouse_Button_Callbacks.Front(); element != nil; element = element.Next() {
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


