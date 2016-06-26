package bl

var g_pluginByName map[string] PlugIn

func init() {
	g_pluginByName = make(map[string] PlugIn)
}

type PlugIn interface {
	Name() string
	Init()
	Tick()
	Uninit()
	OnNodeAdded(node *Node)
	OnNodeRemoved(node *Node)

	GetState() interface{}
	On(cb func(interface{}))
	On2(cb func(interface{}), start func(interface{}), end func(interface{}))
}

func Plugin(p PlugIn) {

	if _, has := g_pluginByName[p.Name()]; !has {
		p.Init()

		g_pluginByName[p.Name()] = p
	}
}

func Use(pluginName string) {
	On(pluginName, nil)
}

func GetPluginState(pluginName string) interface{} {
	plugin, ok := g_pluginByName[pluginName]

	if !ok {
		panic( "GetPluginState() cannot find unregistered plugin named " + pluginName)
	}

	return plugin.GetState()
}

func On(pluginName string, cb func(interface{})) {
	plugin, ok := g_pluginByName[pluginName]

	if !ok {
		panic( "On() cannot find unregistered plugin named " + pluginName)
	}

	plugin.On(cb)
}

func On2(pluginName string, cb func(interface{}), start func(interface{}), end func(interface{})) {
	plugin, ok := g_pluginByName[pluginName]

	if !ok {
		panic( "On2() cannot find unregistered plugin named " + pluginName)
	}

	plugin.On2(cb, start, end)
}
