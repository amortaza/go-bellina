package bl

import (
	"github.com/amortaza/go-bellina/funclist"
)

var g_LifeCycle_Inits = funclist.New()
var g_LifeCycle_Uninits = funclist.New()

var g_LifeCycle_AfterUser_Ticks_ShortTerm = funclist.New()
var g_LifeCycle_AfterUser_Ticks_LongTerm = funclist.New()

// Note: g_LifeCycle_BeforeUser_Ticks_ShortTerm does not make sense.
// ShortTerm can ONLY exist after user ticks!!
var g_LifeCycle_BeforeUser_Ticks_LongTerm = funclist.New()

func Register_LifeCycle_After_UserTick_ShortTerm(cb func()) {
	g_LifeCycle_AfterUser_Ticks_ShortTerm.Add(cb)
}

func Register_LifeCycle_After_UserTick_LongTerm(cb func()) {
	g_LifeCycle_AfterUser_Ticks_LongTerm.Add(cb)
}

func Register_LifeCycle_Before_UserTick_LongTerm(cb func()) {
	g_LifeCycle_BeforeUser_Ticks_LongTerm.Add(cb)
}

func Register_LifeCycle_Init(cb func()) {
	g_LifeCycle_Inits.Add(cb)
}

func Register_LifeCycle_Uninit(cb func()) {
	g_LifeCycle_Uninits.Add(cb)
}
