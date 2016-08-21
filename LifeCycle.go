package bl

import "container/list"

var g_shortTerm_LifeCycleTicks *list.List

func RegisterShortTerm_LifeCycleTick(cb func()) {
	g_shortTerm_LifeCycleTicks.PushBack(cb)
}
