package bl

import (
	"github.com/amortaza/go-hal"
)

var g_user_tick func()
var g_user_init func()
var g_user_uninit func()

var g_graphics hal.Graphics

var Hal hal.HAL

func bl_onAfterGL() {

	g_graphics = Hal.GetGraphics()

	init_bl()

	debug("     (+) Bellina User Init Callback", "sys")

	if g_user_init != nil {
		g_user_init()
	}
}

func bl_onBeforeDeleteWindow() {

	debug("     (-) Calling User Uninit Callback", "sys")

	if g_user_uninit != nil {
		g_user_uninit()
	}

	uninit_bl()
}

func Start(
		hal hal.HAL,

		title string,

		left, top, width, height int,

		user_init func(),
		user_tick func(),
		user_uninit func()) {

	g_user_tick = user_tick
	g_user_init = user_init
	g_user_uninit = user_uninit

	Hal = hal

	hal.Start(
			title,
			left, top, width, height,

			bl_onAfterGL,
			bl_onLoop,
			bl_onBeforeDeleteWindow,

			io_onWindowResize,
			io_onMouseMove,
			io_onMouseButton,
			io_onKey)
}
