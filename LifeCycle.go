package bl

import "container/list"

var g_LifeCycle_Inits = list.New()
var g_LifeCycle_Uninits = list.New()

var g_LifeCycle_AfterUser_Ticks_ShortTerm = list.New()
var g_LifeCycle_AfterUser_Ticks_LongTerm = list.New()

// Note: g_LifeCycle_BeforeUser_Ticks_ShortTerm does not make sense.
// ShortTerm can ONLY exist after user ticks!!
var g_LifeCycle_BeforeUser_Ticks_LongTerm = list.New()

func Register_LifeCycle_After_UserTick_ShortTerm(cb func()) {
	g_LifeCycle_AfterUser_Ticks_ShortTerm.PushBack(cb)
}

func Register_LifeCycle_After_UserTick_LongTerm(cb func()) {
	g_LifeCycle_AfterUser_Ticks_LongTerm.PushBack(cb)
}

func Register_LifeCycle_Before_UserTick_LongTerm(cb func()) {
	g_LifeCycle_BeforeUser_Ticks_LongTerm.PushBack(cb)
}

func Register_LifeCycle_Init(cb func()) {
	g_LifeCycle_Inits.PushBack(cb)
}

func Register_LifeCycle_Uninit(cb func()) {
	g_LifeCycle_Uninits.PushBack(cb)
}

func callAllCallbacks( callbacks *list.List) {

	if callbacks == nil {
		return
	}

	for e := callbacks.Front(); e != nil; e = e.Next() {
		cb := e.Value.(func())
		cb()
	}
}