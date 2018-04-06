package bl

import (
	"container/list"
)

var g_buffersFilled = 0

func bl_onLoop() {

	debug("-------------------- Starting loop", "loop")

	fps()

	// store the last frame
	g_lastFrame_nodeById = g_nodeById

	// Clear Nodes
	g_nodeById = make(map[string] *Node)
	Root_Node = nil

	// Clear Short Term
	g_shortTerm_callbacksByEventType = make(map[string] *list.List)

	g_LifeCycle_AfterUser_Ticks_ShortTerm = list.New()

	g_nodes_are_immutable = false

	// long term ticks
	callAllCallbacks(g_LifeCycle_BeforeUser_Ticks_LongTerm)

	g_user_tick()

	// resize root
	resizeRoot()

	// long term ticks
	callAllCallbacks(g_LifeCycle_AfterUser_Ticks_LongTerm)

	// short term ticks
	callAllCallbacks(g_LifeCycle_AfterUser_Ticks_ShortTerm)

	if Root_Node != nil {

		stabilize(Root_Node)

		syncFromShadow(Root_Node)

		setDirty_New_and_Changed_Nodes(Root_Node)
		setDirty_Removed_Nodes_and_GarbageCollect()

		if Root_Node.Dirty {
			g_buffersFilled = 0
		}

		// on some machines without filling at least 2 buffers, the screen flashes
		// I do not know what is causing this - but this is the workaround
		if g_buffersFilled < 2 {

			g_buffersFilled++
			render()
		}
	}

	if g_nodeStack.Size > 0 {
		panic("Node stack memory leak")
	}

	g_nodes_are_immutable = true

	debug("<<<<<<<<<<<<<<<< exiting loop", "loop")
}

func render() {

	w, h := Hal.GetWindowDim()
	g_graphics.PushView(w, h)

	g_graphics.Clear(.3, .3, .3)

	canvas := renderCanvas(Root_Node)

	canvas.Paint(false, 0, 0, four_ones_float32)

	g_graphics.PopView()
}

func stabilize(node *Node) {

	callAllCallbacks(node.stabilize_funcs_pre_kids)

	for k := node.Kids.Front(); k != nil; k = k.Next() {
		kid := k.Value.(*Node)
		stabilize(kid)
	}

	callAllCallbacks(node.stabilize_funcs_post_kids)
}

func resizeRoot() {

	if Root_Node == nil {
		return
	}

	shadow := EnsureShadowById("ROOT")

	shadow.Width = Window_Width
	shadow.Height = Window_Height
}