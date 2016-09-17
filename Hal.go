package bl

type HAL interface {

	Start(	width, height int,
		title string,
		onAfterGL, onLoop, onBeforeDelete func(),
		onResize, onMouseMove func(int,int),
		onMouseButton func(MouseButton, ButtonAction),
		onKey func(KeyboardKey, ButtonAction, bool, bool, bool))

	GetWindowDim()(width, height int)

	GetMousePos()(x,y int)

	SetMouseCursor(cursor MouseCursor)
}

