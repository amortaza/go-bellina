package bl

import (
	"xel"
	"bellina/core"
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
	canvas.Paint(false, 64, 64, nil)
	canvas.Free()

	g4.PopView()
}

func Start(tick func()) {

	g_tick = tick

	xel.Init(800, 600)

	xel.SetCallbacks(onAfterGL, onLoop, onBeforeDelete, core.IO_onResize, core.IO_onMouseMove, core.IO_onMouseButton, core.IO_onKey)

	xel.Loop()

	xel.Uninit()
}
