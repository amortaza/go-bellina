package bl

type ShadowNode struct {
	Id                                 string
	ParentID                           string

	Left, Top, Width, Height           int

	_backingNode *Node // this is the node that the shadow is backing!!
}

func NewShadowNode(node *Node) *ShadowNode {
	var parentId string

	if node.Parent != nil {
		parentId = node.Parent.Id
	}

	shadow := &ShadowNode{
		node.Id,
		parentId,
		node.Left, node.Top,
		node.Width, node.Height,
		nil,
	}

	return shadow
}