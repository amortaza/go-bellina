package bl

var g_lastFrameNodes map[string] *Node

func init() {
	g_lastFrameNodes = make(map[string] *Node)
}

func setDirty_IncludeKids(node *Node) *Node {

	lastFrameNode, ok := g_lastFrameNodes[node.Id]
	
	if !ok {

		lastFrameNode = &Node{}
		g_lastFrameNodes[node.Id] = lastFrameNode
		
		lastFrameNode.Left = node.Left
		lastFrameNode.Top = node.Top
		lastFrameNode.Width = node.Width
		lastFrameNode.Height = node.Height
		
		node.Dirty = true
		
	} else if lastFrameNode.Width != node.Width || lastFrameNode.Height != node.Height {
		
		node.Dirty = true

		lastFrameNode.Width = node.Width
		lastFrameNode.Height = node.Height
	}
	
	for kide := node.Kids.Front(); kide != nil; kide = kide.Next() {

		kid := kide.Value.(*Node)
		
		lastFrameKid := setDirty_IncludeKids(kid)

		if kid.Dirty {
			node.Dirty = true
		}

		if lastFrameKid.Left != kid.Left || lastFrameKid.Top != kid.Top {

			node.Dirty = true

			lastFrameKid.Left = kid.Left
			lastFrameKid.Top = kid.Top
		}
	}

	return lastFrameNode
}
