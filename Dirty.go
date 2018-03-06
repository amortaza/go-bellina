package bl

var g_lastFrame_nodeById map[string] *Node

func setDirty_IncludeKids(node *Node) *Node {

	lastFrameNode, ok := g_lastFrame_nodeById[node.Id]
	
	if !ok {

		node.Dirty = true
		
	} else if lastFrameNode.Width != node.Width || lastFrameNode.Height != node.Height {
		
		node.Dirty = true
	}
	
	for kide := node.Kids.Front(); kide != nil; kide = kide.Next() {

		kid := kide.Value.(*Node)
		
		lastFrameKid := setDirty_IncludeKids(kid)

		if kid.Dirty {
			node.Dirty = true
		}

		if lastFrameKid.Left != kid.Left || lastFrameKid.Top != kid.Top {

			node.Dirty = true
		}
	}

	return lastFrameNode
}
