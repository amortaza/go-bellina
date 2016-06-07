package bl

import (
	"g4"
	"bellina/event"
	"fmt"
	"xel"
)

func Init() {
	g4.Init()

	g_textureByPartialName = make(map[string] *g4.Texture)

	event.Register(event.Key_Event_Type, func(event event.Event) {

	})

	event.Register(event.Mouse_Move_Event_Type, func(event event.Event) {
	})

	event.Register(event.Mouse_Button_Event_Type, func(event event.Event) {
	})

	event.Register(event.Resize_Event_Type, func(event event.Event) {
		fmt.Println("resize")
	})

	// initial resize fire
	resizeEvent := event.NewResizeEvent(xel.WinWidth, xel.WinHeight)
	event.Fire(resizeEvent)
}

func Uninit() {
	// free texture map images
	for _, value := range g_textureByPartialName {
		value.Free()
	}

	g4.Uninit()
}

