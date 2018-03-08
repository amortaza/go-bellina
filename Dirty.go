package bl

var g_lastFrame_nodeById map[string] *Node

func setDirty_IncludeKids(node *Node) *Node {

	lastFrameNode, ok := g_lastFrame_nodeById[node.Id]

	if !ok {

		node.Dirty = true
		
	} else {

		if lastFrameNode.width != node.width || lastFrameNode.height != node.height {
			node.Dirty = true
		}

		if node.Parent != nil {

			if lastFrameNode.left != node.left || lastFrameNode.top != node.top {
				node.Parent.Dirty = true
			}
		}
	}
	
	for kide := node.Kids.Front(); kide != nil; kide = kide.Next() {

		kid := kide.Value.(*Node)
		
		lastFrameKid := setDirty_IncludeKids(kid)

		if kid.Dirty {
			node.Dirty = true
		}

		if lastFrameKid != nil && (lastFrameKid.left != kid.left || lastFrameKid.top != kid.top) {

			node.Dirty = true
		}
	}

	return lastFrameNode
}
