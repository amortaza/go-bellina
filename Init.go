package bl

import "fmt"

func init_bl() {

	// initial resize fire
	fmt.Println("fire initial Window Resize event")
	resizeEvent := NewWindowResizeEvent(Hal.GetWindowDim())
	FireEvent(resizeEvent)

	// life cycle init
	fmt.Println("(+) Call all Life-Cycle Init methods...")
	callAllCallbacks(g_LifeCycle_Inits)
}

func uninit_bl() {

	// life cycle uninit
	fmt.Println("(-) Call all Life-Cycle Uninit methods...")
	callAllCallbacks(g_LifeCycle_Uninits)

	freeCanvases()
}

