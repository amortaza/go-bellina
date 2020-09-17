package bl

var g_lastFrame_nodeById map[string] *Node

func setDirty_New_and_Changed_Nodes(node *Node) *Node {
	lastFrameNode, ok := g_lastFrame_nodeById[node.Id]

	if !ok {

		// making parents dirty is done in child loop below
		node.Dirty = true

	} else if lastFrameNode.width != node.width || lastFrameNode.height != node.height {

		// making parents dirty is done in child loop below
		node.Dirty = true
	}

	for e := node.Kids.Front(); e != nil; e = e.Next() {
		kid := e.Value.(*Node)
		
		lastFrameKid := setDirty_New_and_Changed_Nodes(kid)

		if kid.Dirty {
			node.Dirty = true
		}

		if lastFrameKid != nil && (lastFrameKid.left != kid.left || lastFrameKid.top != kid.top) {
			node.Dirty = true
		}
	}

	return lastFrameNode
}

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
