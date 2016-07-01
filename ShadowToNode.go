package bl

import "fmt"

func (s *ShadowNode) Dim(width, height int) *ShadowNode {
	s.Width = width
	s.Height = height

	s._backingNode.Width = s.Width
	s._backingNode.Height = s.Height

	//fmt.Println("set shadow width ", s.Id, s.Width)

	return s
}

func (s *ShadowNode) Dim_to_Node() *ShadowNode {
	//fmt.Println("shadow width ", s.Id, s.Width)
	s._backingNode.Width = s.Width
	s._backingNode.Height = s.Height

	return s
}

func (s *ShadowNode) PosLeft(left int) *ShadowNode {
	fmt.Println("wth")
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

func (s *ShadowNode) Color1(r,g,b float32) *ShadowNode {
	s.Red1, s.Green1, s.Blue1 = r, g, b

	s._backingNode.Red1, s._backingNode.Green1, s._backingNode.Blue1 = r, g, b

	return s
}

func (s *ShadowNode) Color1_to_Node() *ShadowNode {
	s._backingNode.Red1, s._backingNode.Green1, s._backingNode.Blue1 = s.Red1, s.Green1, s.Blue1

	return s
}

func (s *ShadowNode) FontColor(r,g,b float32) *ShadowNode {
	s.FontRed, s.FontGreen, s.FontBlue = r, g, b

	s._backingNode.FontRed, s._backingNode.FontGreen, s._backingNode.FontBlue = s.FontRed, s.FontGreen, s.FontBlue

	return s
}

func (s *ShadowNode) FontColor_to_Node() *ShadowNode {
	s._backingNode.FontRed, s._backingNode.FontGreen, s._backingNode.FontBlue = s.FontRed, s.FontGreen, s.FontBlue

	return s
}


