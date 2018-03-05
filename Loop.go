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
	for e := g_LifeCycle_BeforeUser_Ticks_LongTerm.Front(); e != nil; e = e.Next() {
		cb := e.Value.(func())
		cb()
	}

	g_user_tick()

	// long term ticks
	for e := g_LifeCycle_AfterUser_Ticks_LongTerm.Front(); e != nil; e = e.Next() {
		cb := e.Value.(func())
		cb()
	}

	// short term ticks
	for e := g_LifeCycle_AfterUser_Ticks_ShortTerm.Front(); e != nil; e = e.Next() {
		cb := e.Value.(func())
		cb()
	}

	// todo
	g_graphics.Clear(.3, .3, .3)

	w, h := Hal.GetWindowDim()
	g_graphics.PushView(w, h)

	debug("**** CANVAS")

	if Root_Node != nil {
		Stabilize(Root_Node)

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

