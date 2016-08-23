package bl

import "container/list"

var g_LifeCycle_Inits = list.New()
var g_LifeCycle_Uninits = list.New()

var g_LifeCycle_AfterUser_Ticks_ShortTerm *list.List
var g_LifeCycle_AfterUser_Ticks = list.New()

var g_LifeCycle_BeforeUser_Ticks_ShortTerm *list.List
var g_LifeCycle_BeforeUser_Ticks = list.New()

func Register_LifeCycle_AfterUser_Tick_ShortTerm(cb func()) {
	g_LifeCycle_AfterUser_Ticks_ShortTerm.PushBack(cb)
}

func Register_LifeCycle_BeforeUser_Tick_ShortTerm(cb func()) {
	g_LifeCycle_BeforeUser_Ticks_ShortTerm.PushBack(cb)
}

func Register_LifeCycle_AfterUser_Tick(cb func()) {
	g_LifeCycle_AfterUser_Ticks.PushBack(cb)
}

func Register_LifeCycle_BeforeUser_Tick(cb func()) {
	g_LifeCycle_BeforeUser_Ticks.PushBack(cb)
}

func Register_LifeCycle_Init(cb func()) {
	g_LifeCycle_Inits.PushBack(cb)
}

func Register_LifeCycle_Uninit(cb func()) {
	g_LifeCycle_Uninits.PushBack(cb)
}
