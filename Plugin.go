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

	if _, has := g_pluginByName[p.Name()]; !has {
		p.Init()

		g_pluginByName[p.Name()] = p
	}
}

func On(pluginName string, cb func(interface{})) {
	plugin, ok := g_pluginByName[pluginName]

	if !ok {
		panic( "On() cannot find unregistered plugin named " + pluginName)
	}

	plugin.On(cb)
}
