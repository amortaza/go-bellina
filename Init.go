package bl

import (
	"github.com/amortaza/go-g5"
	"container/list"
)

func Init() {
	g5.Init()

	g_shadowNodeByID = make(map[string] *ShadowNode)
	g_textureByPartialName = make(map[string] *g5.Texture)

	g_pluginTicks = list.New()

	// initial resize fire
	resizeEvent := NewResizeEvent(g_hal.GetWindowDim())
	FireEvent(resizeEvent)
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

	g5.Uninit()
}

