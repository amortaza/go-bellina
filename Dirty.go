package bl

var g_lastFrame_nodeById map[string] *Node

func setDirty_New_and_Changed_Nodes(node *Node) *Node {

	lastFrameNode, ok := g_lastFrame_nodeById[node.Id]

	if !ok {

		// making parents dirty is done in child loop
		node.Dirty = true

	} else {

		if lastFrameNode.width != node.width || lastFrameNode.height != node.height {

			// making parents dirty is done in child loop
			node.Dirty = true
		}
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
