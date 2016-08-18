package bl

import "container/list"

var g_pluginByName map[string] PlugIn
var g_plugins_inOrder *list.List
var g_pluginTicks *list.List

func init() {
	g_pluginByName = make(map[string] PlugIn)
	g_plugins_inOrder = list.New()
}

type PlugIn interface {
	Name() string
	Init()
	Tick()
	Reset_ShortTerm()
	Uninit()
	OnNodeAdded(node *Node)
	OnNodeRemoved(node *Node)
}

func Plugin(p PlugIn) {

	if _, has := g_pluginByName[p.Name()]; !has {
		p.Init()

		g_pluginByName[p.Name()] = p
		g_plugins_inOrder.PushBack(p)
	}
}

func AddPluginOnTick(cb func()) {
	g_pluginTicks.PushBack(cb)
}
