package bl

import (
	"container/list"
	"github.com/amortaza/go-g5"
)

func onLoop() {

	// Clear Nodes
	g_nodeById = make(map[string] *Node)

	// Clear Plugin On Tick
	g_pluginTicks = list.New()

	// Clear Short Term
	g_shortTerm_callbacksByEventType = make(map[string] *list.List)

	for e := g_plugins_inOrder.Front(); e != nil; e = e.Next() {
		plugin := e.Value.(PlugIn)
		plugin.Reset_ShortTerm()
	}

	g_tick()

	for e := g_plugins_inOrder.Front(); e != nil; e = e.Next() {
		plugin := e.Value.(PlugIn)
		plugin.Tick()
	}

	for e := g_pluginTicks.Front(); e != nil; e = e.Next() {
		cb := e.Value.(func())
		cb()
	}

	g5.Clear(.3,.3,.3,1)

	w, h := g_hal.GetWindowDim()
	g5.PushView(w, h)

	canvas := renderCanvas(Root_Node)

	if canvas != nil {
		canvas.Paint(false, 0, 0, g5.FourOnesFloat32)
		canvas.Free()
	}

	g5.PopView()

	// #ifdef nonprod
	if g_nodeStack.Size > 0 {
		panic("Node stack memory leak")
	}
}

