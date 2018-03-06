package bl

func garbageCollectDeletedNodes() {

	for nodeId, node := range g_lastFrame_nodeById {

		_, ok := g_nodeById[nodeId]

		if !ok {
			freeNode(node)
		}
	}
}

func freeNode(node *Node) {

	debug("      Freeing node " + node.Id, "gc")

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
