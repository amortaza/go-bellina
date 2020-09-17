package funclist

import "container/list"

type FuncList struct {
	funcList *list.List
}

func New() *FuncList {
	return &FuncList{funcList: list.New()}
}

func (flist *FuncList) Add(f func()) {
	if flist.funcList == nil {
		flist.funcList = list.New()
	}

	flist.funcList.PushBack(f)
}

func (flist *FuncList) CallAll() {
	if flist.funcList == nil {
		return
	}

	for e := flist.funcList.Front(); e != nil; e = e.Next() {
		cb := e.Value.(func())
		cb()
	}
}
