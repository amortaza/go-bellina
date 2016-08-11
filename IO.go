package bl

import (
	"github.com/amortaza/go-bellina/event"
	"github.com/amortaza/go-bellina/constants"
)

func IO_onKey(key bl.KeyboardKey, action bl.ButtonAction, alt, ctrl, shift bool) {
	keyEvent := NewKeyEvent(key, action, alt, ctrl, shift)

	event.Fire(keyEvent)
}

func IO_onMouseMove(x,y int) {
	//x /= DevicePixelRatio
	//y /= DevicePixelRatio

	Mouse_X_Scaled, Mouse_Y_Scaled = x, y

	node := GetNodeAt_VisibleToEvents(x, y)

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

	event.Fire(e)
}

func IO_onMouseButton(button bl.MouseButton, action bl.ButtonAction) {
	node := GetNodeAt_VisibleToEvents(Mouse_X_Scaled, Mouse_Y_Scaled)

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

	event.Fire(e)
}

func IO_onResize(width, height int) {
	resizeEvent := event.NewResizeEvent(width, height)

	event.Fire(resizeEvent)
}


