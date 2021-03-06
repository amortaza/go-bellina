package bl

import (
	"container/list"
	"github.com/amortaza/go-bellina/debug"
	"github.com/amortaza/go-bellina/funclist"
)

var g_buffersFilled = 0
var g_fps = debug.NewFPS()

func bl_onLoop() {
	debug.Log("-------------------- Starting loop", debug.Loop)

	g_fps.Tick()

	// store the last frame
	g_lastFrame_nodeById = g_nodeById

	// Clear Nodes
	debug.Log("-------------------- Clearning Nodes", debug.Loop)
	g_nodeById = make(map[string] *Node)
	Root_Node = nil

	// Clear Short Term
	g_shortTerm_callbacksByEventType = make(map[string] *list.List)
	g_LifeCycle_AfterUser_Ticks_ShortTerm = funclist.New()

	g_nodes_are_immutable = false

	// long term ticks
	g_LifeCycle_BeforeUser_Ticks_LongTerm.CallAll()

	debug.Log("-------------------- Starting to build nodes", debug.Loop)
	g_user_tick()
	debug.Log("-------------------- Finished building nodes", debug.Loop)

	//
	sizeRootShadowToWindow()

	// long term ticks
	g_LifeCycle_AfterUser_Ticks_LongTerm.CallAll()

	// short term ticks
	g_LifeCycle_AfterUser_Ticks_ShortTerm.CallAll()

	if Root_Node != nil {
		stabilize(Root_Node)

		syncFromShadow_Position_and_Dimension(Root_Node)

		mark_new_and_changed_nodes_and_their_parents_dirty(Root_Node)
		garbage_collect_removed_nodes_and_mark_their_parents_dirty()

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

	debug.Log("<<<<<<<<<<<<<<<< exiting loop", debug.Loop)
}

func render() {
	w, h := g_hal.GetWindowDim()
	g_graphics.PushView(w, h)

	g_graphics.Clear(.3, .3, .3)

	canvas := renderCanvas(Root_Node)

	canvas.Paint(true, 0, 0, g_four_ones)

	g_graphics.PopView()
}

func stabilize(node *Node) {
	if node.stabilize_funcs_pre_kids != nil {
		node.stabilize_funcs_pre_kids.CallAll()
	}

	for k := node.Kids.Front(); k != nil; k = k.Next() {
		kid := k.Value.(*Node)
		stabilize(kid)
	}

	if node.stabilize_funcs_post_kids != nil {
		node.stabilize_funcs_post_kids.CallAll()
	}
}

func sizeRootShadowToWindow() {
	if Root_Node == nil {
		return
	}

	shadow := EnsureShadowById("ROOT")

	shadow.Width = Window_Width
	shadow.Height = Window_Height
}