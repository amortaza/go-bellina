package bl

import (
	"container/list"
	"github.com/amortaza/go-g5"
	"fmt"	
)

func fake3() {
    var _ = fmt.Println
}

func onLoop() {
	debug := false

	if debug {
		fmt.Println("START ********************************************************************** ", )
	}

	// Clear Nodes
	g_nodeById = make(map[string] *Node)

	// Clear Short Term
	g_shortTerm_callbacksByEventType = make(map[string] *list.List)

	g_LifeCycle_BeforeUser_Ticks_ShortTerm = list.New()
	g_LifeCycle_AfterUser_Ticks_ShortTerm = list.New()

	// long term ticks
	for e := g_LifeCycle_BeforeUser_Ticks.Front(); e != nil; e = e.Next() {
		cb := e.Value.(func())
		cb()
	}

	// short term ticks
	for e := g_LifeCycle_BeforeUser_Ticks_ShortTerm.Front(); e != nil; e = e.Next() {
		cb := e.Value.(func())
		cb()
	}

	g_tick()

	// long term ticks
	for e := g_LifeCycle_AfterUser_Ticks.Front(); e != nil; e = e.Next() {
		cb := e.Value.(func())
		cb()
	}

	// stabilize
	for i := 0; i <= g_root_depth + 1; i++ {
		for e := g_funcs.Front(); e != nil; e = e.Next() {
			cb := e.Value.(func())
			cb()
		}
	}

	g_funcs.Init()
	g_root_depth = 0

	g5.Clear(.3,.3,.3,1)

	w, h := g_hal.GetWindowDim()
	g5.PushView(w, h)

	if debug {
		fmt.Println("**** CANVAS")
	}

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

	if debug {
		fmt.Println("END ********************************************************************** ", )
	}
}

