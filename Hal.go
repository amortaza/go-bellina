package bl

type Canvas interface {
	Begin()
	End()
	GetWidth() int
	GetHeight() int
	Clear(red, green, blue float32)
	Paint(seeThru bool, left, top int, alphas []float32)
	Free()
}

type Graphics interface {
	Clear(red, green, blue, alpha float32)
	PushView(width, height int)
	PopView()
	NewCanvas(width, height int) Canvas
}

type HAL interface {

	Start(	width, height int,
		title string,
		onAfterGL, onLoop, onBeforeDelete func(),
		onWindowResize, onMouseMove func(int,int),
		onMouseButton func(MouseButton, ButtonAction),
		onKey func(KeyboardKey, ButtonAction, bool, bool, bool))

	GetWindowDim()(width, height int)

	GetMousePos()(x,y int)

	SetMouseCursor(cursor MouseCursor)

	GetGraphics() Graphics
}

