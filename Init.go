package bl

import "github.com/amortaza/go-bellina/debug"

func init_bl() {

	// initial resize fire
	debug.Log("     Fire initial Window Resize event", debug.System)
	resizeEvent := NewWindowResizeEvent(g_hal.GetWindowDim())
	FireEvent(resizeEvent)

	// life cycle init
	debug.Log("     (+) Call all Life-Cycle Init methods...", debug.System)
	g_LifeCycle_Inits.CallAll()
}

func uninit_bl() {

	// life cycle uninit
	debug.Log("     (-) Call all Life-Cycle Uninit methods...", debug.System)
	g_LifeCycle_Uninits.CallAll()

	freeCanvases()
}

