package bl

import (
	"xel"
	"g4"
)

var g_tick func()
var g_init func()
var g_uninit func()

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

func onLoop() {

	g_tick()

	g4.Clear(.4,.6,.6,1)

	g4.PushView(xel.WinWidth, xel.WinHeight)

	canvas := renderCanvas(Root_Node)
	canvas.Paint(false, Root_Node.Left, Root_Node.Top, nil) // also modify spatial
	canvas.Free()

	g4.PopView()

	// #ifdef nonprod
	if g_nodeStack.Size > 0 {
		panic("Node stack memory leak")
	}
}

func Start(width, height int32, title string, init func(), tick func(), uninit func()) {

	g_tick = tick
	g_init = init
	g_uninit = uninit

	xel.Init(width, height)

	xel.SetCallbacks(onAfterGL, onLoop, onBeforeDelete, IO_onResize, IO_onMouseMove, IO_onMouseButton, IO_onKey)

	xel.Loop()

	xel.Uninit()
}
