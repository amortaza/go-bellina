package bl

import "container/list"

func bl_onLoop() {

	// todo remove for PROD
	fps()

	debug("START ********************************************************************** " )

	// Clear Nodes
	g_nodeById = make(map[string] *Node)

	// Clear Short Term
	g_shortTerm_callbacksByEventType = make(map[string] *list.List)

	g_LifeCycle_AfterUser_Ticks_ShortTerm = list.New()

	// long term ticks
	callAllCallbacks(g_LifeCycle_BeforeUser_Ticks_LongTerm)

	g_user_tick()

	// long term ticks
	callAllCallbacks(g_LifeCycle_AfterUser_Ticks_LongTerm)

	// short term ticks
	callAllCallbacks(g_LifeCycle_AfterUser_Ticks_ShortTerm)

	// todo
	g_graphics.Clear(.3, .3, .3)

	w, h := Hal.GetWindowDim()
	g_graphics.PushView(w, h)

	debug("**** CANVAS")

	if Root_Node != nil {

		stabilize(Root_Node)

		setDirty_IncludeKids(Root_Node)

		canvas := renderCanvas(Root_Node)

		canvas.Paint(false, 0, 0, FourOnesFloat32)
	}

	g_graphics.PopView()

	// todo prod
	if g_nodeStack.Size > 0 {
		panic("Node stack memory leak")
	}

	debug("END ********************************************************************** ")
}

func stabilize(node *Node) {

	callAllCallbacks(node.stabilize_funcs_pre_kids)

	for k := node.Kids.Front(); k != nil; k = k.Next() {
		kid := k.Value.(*Node)
		stabilize(kid)
	}

	callAllCallbacks(node.stabilize_funcs_post_kids)
}
