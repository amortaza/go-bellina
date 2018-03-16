package bl

import (
	"container/list"
)

func init() {

	g_onFreeNodeCallbacks = make(map[string] *list.List)
}

var g_onFreeNodeCallbacks map[string] *list.List

func mark_this_and_parents_dirty(node *Node) {

	node.Dirty = true

	var p = node.Parent

	for true {

		if p == nil {
			return
		}

		if p.Dirty {
			return
		}

		p.Dirty = true

		p = p.Parent
	}
}

func setDirty_Removed_Nodes_and_GarbageCollect() {

	// find deleted nodes
	for nodeId, node := range g_lastFrame_nodeById {

		_, ok := g_nodeById[nodeId]

		if !ok {

			if node.Parent != nil {

				// everything so far is from 'previous' frame, get into 'this' frame
				parent, ok := g_nodeById[node.Parent.Id]

				if ok {
					mark_this_and_parents_dirty(parent)
				}
			}

			freeNode(node)
		}
	}
}

func freeNode(node *Node) {

	debug("      Freeing node " + node.Id, "gc")

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

		debug("      (-) Freeing canvas for " + nodeId, "gc")

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