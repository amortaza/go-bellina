package bl

import (
	"github.com/amortaza/go-xel"
	"github.com/amortaza/go-g4"
	"container/list"
	"github.com/amortaza/go-bellina/event"
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

func Start(width, height int, title string, init func(), tick func(), uninit func()) {

	g_tick = tick
	g_init = init
	g_uninit = uninit

	xel.Init(width, height)

	xel.SetCallbacks(onAfterGL, onLoop, onBeforeDelete, IO_onResize, IO_onMouseMove, IO_onMouseButton, IO_onKey)

	xel.Loop()

	xel.Uninit()
}
