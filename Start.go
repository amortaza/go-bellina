package bl

import "fmt"

var g_tick func()
var g_init func()
var g_uninit func()

type ButtonAction int
type MouseButton int
type KeyboardKey int
type MouseCursor int

var Hal HAL

func onAfterGL() {
	init_bl()

	fmt.Println("+ user init callback")
	if g_init != nil {
		g_init()
	}


}

func onBeforeDelete() {
	fmt.Println("- user uninit callback")
	if g_uninit != nil {
		g_uninit()
	}

	uninit_bl()
}

func Start(	hal HAL,
		width, height int,
		title string,
		init func(),
		tick func(),
		uninit func()) {

	g_tick = tick
	g_init = init
	g_uninit = uninit

	Hal = hal

	hal.Start(	width, height,
			title,
			onAfterGL,
			onLoop,
			onBeforeDelete,
			io_onWindowResize,
			io_onMouseMove,
			io_onMouseButton,
			io_onKey)
}
