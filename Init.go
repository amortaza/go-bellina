package bl

import (
	"github.com/amortaza/go-g4"
	"github.com/amortaza/go-bellina/event"
	"github.com/amortaza/go-xel"
	"container/list"
)

func Init() {
	g4.Init()

	g_shadowNodeByID = make(map[string] *ShadowNode)
	g_textureByPartialName = make(map[string] *g4.Texture)

	g_pluginTicks = list.New()

	g_pluginParamsNodeId_int = make(map[string] (map[string] int))
	g_pluginParamsNodeId_string = make(map[string] (map[string] string))
	g_pluginParamsNodeId_func = make(map[string] (map[string] func()))

	// initial resize fire
	resizeEvent := event.NewResizeEvent(xel.WinWidth, xel.WinHeight)
	event.Fire(resizeEvent)
}

func Uninit() {
	// free plugins
	for _, plugin := range g_pluginByName {

		plugin.Uninit()
	}

	// free texture map images
	for _, value := range g_textureByPartialName {
		value.Free()
	}

	g4.Uninit()
}

