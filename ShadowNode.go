package bl

type ShadowNode struct {

	Id string

	Left, Top, Width, Height int

	BackingNode *Node // this is the node that the shadow is backing!!
}

func newShadowNode(node *Node) *ShadowNode {

	shadow := &ShadowNode{
		node.Id,
		node.left, node.top,
		node.width, node.height,
		nil, // this will get overwritten anyways, do not set to node
	}

	return shadow
}




