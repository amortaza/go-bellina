package bl

var g_tick func()
var g_init func()
var g_uninit func()

var g_hal Hal

type Hal interface {

	Start(	width, height int,
		title string,
		onAfterGL, onLoop, onBeforeDelete func(),
		onResize, onMouseMove func(int,int),
		onMouseButton func(MouseButton, ButtonAction),
		onKey func(KeyboardKey, ButtonAction, bool, bool, bool))

	GetWindowDim()(width, height int)
}

func onAfterGL() {
	Init()

	if g_init != nil {
		g_init()
	}
}

func onBeforeDelete() {
	if g_uninit != nil {
		g_uninit()
	}

	Uninit()
}

func Start(	hal Hal,
		width, height int,
		title string,
		init func(),
		tick func(),
		uninit func()) {

	g_tick = tick
	g_init = init
	g_uninit = uninit

	g_hal = hal

	hal.Start(	width, height,
			title,
			onAfterGL,
			onLoop,
			onBeforeDelete,
			IO_onResize,
			IO_onMouseMove,
			IO_onMouseButton,
			IO_onKey)
}
