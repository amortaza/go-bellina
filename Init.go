package bl

import (
	"github.com/amortaza/go-g5"
	"fmt"
)

func init_bl() {
	fmt.Println("(+) g5.Init()")
	g5.Init()

	// initial resize fire
	fmt.Println("(+) initial Resize fire event")
	resizeEvent := NewWindowResizeEvent(Hal.GetWindowDim())
	FireEvent(resizeEvent)

	// life cycle init
	fmt.Println("(+) life cycle init callbacks...")
	for e := g_LifeCycle_Inits.Front(); e != nil; e = e.Next() {
		cb := e.Value.(func())
		cb()
	}
}

func uninit_bl() {
	// life cycle uninit
	fmt.Println("(-) life cycle uninit callbacks...")
	for e := g_LifeCycle_Uninits.Front(); e != nil; e = e.Next() {
		cb := e.Value.(func())
		cb()
	}

	fmt.Println("(-) g5.Uninit()")
	g5.Uninit()
}

