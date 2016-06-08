package bl

import "container/list"

var g_plugins *list.List

func init() {
	g_plugins = list.New()
}

type PlugIn interface {
	Name() string
	Init()
	Uninit()
}

func Plugin(p PlugIn) {
	p.Init()

	g_plugins.PushBack(p)
}
