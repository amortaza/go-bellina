package bl

import (
	"container/list"
)

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

	// long term ticks
	callAllCallbacks(g_LifeCycle_AfterUser_Ticks_LongTerm)

	// short term ticks
	callAllCallbacks(g_LifeCycle_AfterUser_Ticks_ShortTerm)

	g_graphics.Clear(.3, .3, .3)

	w, h := Hal.GetWindowDim()
	g_graphics.PushView(w, h)

	if Root_Node != nil {

		stabilize(Root_Node)

		setDirty_IncludeKids(Root_Node)

		canvas := renderCanvas(Root_Node)

		// todo: root is being painted all the time
		canvas.Paint(false, 0, 0, FourOnesFloat32)
	}

	g_graphics.PopView()

	if g_nodeStack.Size > 0 {
		panic("Node stack memory leak")
	}

	g_nodes_are_immutable = true

	garbageCollectDeletedNodes()

	debug("<<<<<<<<<<<<<<<< exiting loop", "loop")
}

func stabilize(node *Node) {

	callAllCallbacks(node.stabilize_funcs_pre_kids)

	for k := node.Kids.Front(); k != nil; k = k.Next() {
		kid := k.Value.(*Node)
		stabilize(kid)
	}
}
