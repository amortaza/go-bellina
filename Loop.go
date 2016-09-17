package bl

import (
	"container/list"
	"github.com/amortaza/go-g5"
)

func onLoop() {
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

	g_tick()

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
	g5.Clear(.3,.3,.3,1)

	w, h := Hal.GetWindowDim()
	g5.PushView(w, h)

	debug("**** CANVAS")

	if Root_Node != nil {
		Stabilize(Root_Node)

		setDirty_IncludeKids(Root_Node)

		canvas := renderCanvas(Root_Node)

		canvas.Paint(false, 0, 0, g5.FourOnesFloat32)
	}

	g5.PopView()

	// todo prod
	if g_nodeStack.Size > 0 {
		panic("Node stack memory leak")
	}

	debug("END ********************************************************************** ")
}

