package bl

import (
	"github.com/amortaza/go-g5"
	"container/list"
	"github.com/amortaza/go-bellina/event"
	"github.com/amortaza/go-bellina/constants"
	"github.com/amortaza/go-ux"
)

var g_tick func()
var g_init func()
var g_uninit func()
var g_hal Hal

var DevicePixelRatio = 2

type Hal interface {
	Start(	width, height int,
		title string,
		onAfterGL, onLoop, onBeforeDelete func(),
		onResize, onMouseMove func(int,int),
		onMouseButton func(bl.MouseButton, bl.ButtonAction),
		onKey func(bl.KeyboardKey, bl.ButtonAction, bool, bool, bool))

	GetWindowDim()(width, height int)
}

func onAfterGL() {
	Init()

	ux.Init()

	if g_init != nil {
		g_init()
	}
}

func onBeforeDelete() {
	if g_uninit != nil {
		g_uninit()
	}

	ux.Uninit()

	Uninit()
}

func onLoop() {

	g_nodeByID_Previous = g_nodeByID
	g_nodeByID = make(map[string] *Node)
	g_pluginTicks = list.New()

	event.G_registerShortTermCallbacksByEventType = make(map[string] *list.List)

	for e := g_pluginsInOrder.Front(); e != nil; e = e.Next() {
		plugin := e.Value.(PlugIn)
		plugin.Reset()
	}

	g_tick()

	for e := g_pluginsInOrder.Front(); e != nil; e = e.Next() {
	        plugin := e.Value.(PlugIn)
		plugin.Tick()
	}

	for e := g_pluginTicks.Front(); e != nil; e = e.Next() {
	    	cb := e.Value.(func())
		cb()
	}

	detectDifferences(g_nodeByID_Previous, g_nodeByID)

	g5.Clear(.33,.33,.33,1)

	w, h := g_hal.GetWindowDim()
	g5.PushView(w, h)

	ux.Ctx.BeginFrame(w/DevicePixelRatio, h/DevicePixelRatio, float32(DevicePixelRatio))

	canvas := renderCanvas(Root_Node)
	canvas.Paint(false, 0, 0, g5.FourOnesFloat32)
	canvas.Free()

	ux.Ctx.EndFrame()

	g5.PopView()

	// #ifdef nonprod
	if g_nodeStack.Size > 0 {
		panic("Node stack memory leak")
	}
}

func Start(hal Hal, width, height int, title string, init func(), tick func(), uninit func()) {

	g_tick = tick
	g_init = init
	g_uninit = uninit

	g_hal = hal

	hal.Start(width, height, title, onAfterGL, onLoop, onBeforeDelete, IO_onResize, IO_onMouseMove, IO_onMouseButton, IO_onKey)
}
