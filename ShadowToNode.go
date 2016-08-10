package bl

func (s *ShadowNode) Dim(width, height int) *ShadowNode {
	s.Width = width
	s.Height = height

	s._backingNode.Width = s.Width
	s._backingNode.Height = s.Height

	return s
}

func (s *ShadowNode) Dim_to_Node() *ShadowNode {
	s._backingNode.Width = s.Width
	s._backingNode.Height = s.Height

	return s
}

func (s *ShadowNode) PosLeft(left int) *ShadowNode {
	s.Left = left

	s._backingNode.Left = s.Left

	return s
}

func (s *ShadowNode) PosLeft_to_Node() *ShadowNode {
	s._backingNode.Left = s.Left

	return s
}

func (s *ShadowNode) PosTop(top int) *ShadowNode {
	s.Top = top

	s._backingNode.Top = s.Top

	return s
}

func (s *ShadowNode) PosTop_to_Node() *ShadowNode {
	s._backingNode.Top = s.Top

	return s
}



