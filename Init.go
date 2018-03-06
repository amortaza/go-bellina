package bl

func init_bl() {

	// initial resize fire
	debug("     Fire initial Window Resize event", "sys")

	resizeEvent := NewWindowResizeEvent(Hal.GetWindowDim())
	FireEvent(resizeEvent)

	// life cycle init
	debug("     (+) Call all Life-Cycle Init methods...", "sys")
	callAllCallbacks(g_LifeCycle_Inits)
}

func uninit_bl() {

	// life cycle uninit
	debug("     (-) Call all Life-Cycle Uninit methods...", "sys")
	callAllCallbacks(g_LifeCycle_Uninits)

	freeCanvases()
}

