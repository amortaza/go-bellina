package bl

import (
	"github.com/amortaza/go-g5"
	"fmt"
)

func Init() {
	fmt.Println("(+) g5.Init()")
	g5.Init()

	// initial resize fire
	fmt.Println("(+) initial Resize fire event")
	resizeEvent := NewResizeEvent(g_hal.GetWindowDim())
	FireEvent(resizeEvent)

	// life cycle init
	fmt.Println("(+) life cycle init callbacks...")
	for e := g_LifeCycle_Inits.Front(); e != nil; e = e.Next() {
		cb := e.Value.(func())
		cb()
	}
}

func Uninit() {
	// life cycle uninit
	fmt.Println("(-) life cycle uninit callbacks...")
	for e := g_LifeCycle_Uninits.Front(); e != nil; e = e.Next() {
		cb := e.Value.(func())
		cb()
	}

	fmt.Println("(-) g5.Uninit()")
	g5.Uninit()
}

