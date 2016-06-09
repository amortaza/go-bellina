package bl


var g_pluginByName map[string] PlugIn

func init() {
	g_pluginByName = make(map[string] PlugIn)
}

type PlugIn interface {
	Name() string
	Init()
	Uninit()

	On(cb func(interface{}))
}

func Plugin(p PlugIn) {
	p.Init()

	g_pluginByName[p.Name()] = p
}

func On(pluginName string, cb func(interface{})) {
	plugin := g_pluginByName[pluginName]

	plugin.On(cb)
}
