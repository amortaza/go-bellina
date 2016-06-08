package bl

import (
	"xel"
	"g4"
)

var g_tick func()

func onAfterGL() {
	Init()
}

func onBeforeDelete() {
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
}

func Start(width, height int32, title string, tick func()) {

	g_tick = tick

	xel.Init(width, height)

	xel.SetCallbacks(onAfterGL, onLoop, onBeforeDelete, IO_onResize, IO_onMouseMove, IO_onMouseButton, IO_onKey)

	xel.Loop()

	xel.Uninit()
}
