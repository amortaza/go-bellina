package bl

import "container/list"

var g_LifeCycle_Inits = list.New()
var g_LifeCycle_Uninits = list.New()
var g_LifeCycle_ShortTerm_Ticks *list.List
var g_LifeCycle_LongTerm_Ticks = list.New()

func Register_LifeCycle_ShortTerm_Tick(cb func()) {
	g_LifeCycle_ShortTerm_Ticks.PushBack(cb)
}

func Register_LifeCycle_LongTerm_Tick(cb func()) {
	g_LifeCycle_LongTerm_Ticks.PushBack(cb)
}

func Register_LifeCycle_Init(cb func()) {
	g_LifeCycle_Inits.PushBack(cb)
}

func Register_LifeCycle_Uninit(cb func()) {
	g_LifeCycle_Uninits.PushBack(cb)
}
