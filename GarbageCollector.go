package bl

import (
	"container/list"
	"github.com/amortaza/go-bellina/debug"
)

func init() {
	g_onFreeNodeCallbacks = make(map[string] *list.List)
}

var g_onFreeNodeCallbacks map[string] *list.List

func freeNode(node *Node) {
	debug.Log("      Freeing node " + node.Id, debug.GC)

	freeNode_fromPluginData(node.Id)

	delete(g_shadowNodeById, node.Id)

	canvas, ok := g_canvasById[node.Id]

	if ok {
		canvas.Free()
		delete(g_canvasById, node.Id)
	}
}

func freeCanvases() {
	for nodeId, canvas := range g_canvasById {
		debug.Log("      (-) Freeing canvas for " + nodeId, debug.GC)

		canvas.Free()
	}
}

func freeNode_fromPluginData(nodeId string) {
	callbacks, ok := g_onFreeNodeCallbacks[nodeId]

	if !ok {
		return
	}

	for e := callbacks.Front(); e != nil; e = e.Next() {
		callback := e.Value.(func(string))

		callback(nodeId)
	}
}