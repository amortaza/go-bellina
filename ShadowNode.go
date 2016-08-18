package bl

type ShadowNode struct {
	Id                       string
	ParentId                 string

	Left, Top, Width, Height int

	_backingNode             *Node // this is the node that the shadow is backing!!
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

func (s *ShadowNode) Dim__Self_and_Node(width, height int) *ShadowNode {
	s.Width = width
	s.Height = height

	s._backingNode.Width = s.Width
	s._backingNode.Height = s.Height

	return s
}

func (s *ShadowNode) Dim__Node_Only() *ShadowNode {
	s._backingNode.Width = s.Width
	s._backingNode.Height = s.Height

	return s
}

func (s *ShadowNode) PosLeft__Self_and_Node(left int) *ShadowNode {
	s.Left = left

	s._backingNode.Left = s.Left

	return s
}

func (s *ShadowNode) PosLeft__Node_Only() *ShadowNode {
	s._backingNode.Left = s.Left

	return s
}

func (s *ShadowNode) PosTop__Self_and_Node(top int) *ShadowNode {
	s.Top = top

	s._backingNode.Top = s.Top

	return s
}

func (s *ShadowNode) PosTop__Node_Only() *ShadowNode {
	s._backingNode.Top = s.Top

	return s
}



