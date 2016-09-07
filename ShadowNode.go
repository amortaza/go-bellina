package bl

import "fmt"

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

func hasWidthOwner(s *ShadowNode, owner string) bool {
	if owner == "self" && (s._backingNode.OwnerOfWidth != "" && s._backingNode.OwnerOfWidth != "self") {
		fmt.Println("\"", owner, "\" tried to set Width, but Node \"", s._backingNode.Id, "\" width is already owned by \"", s._backingNode.OwnerOfWidth, "\"")

		return true
	}

	return false
}

func hasHeightOwner(s *ShadowNode, owner string) bool {
	if owner == "self" && (s._backingNode.OwnerOfHeight != "" && s._backingNode.OwnerOfHeight != "self") {
		fmt.Println("\"", owner, "\" tried to set Height, but Node \"", s._backingNode.Id, "\" height is already owned by \"", s._backingNode.OwnerOfHeight, "\"")

		return true
	}

	return false

}

func hasLeftOwner(s *ShadowNode, owner string) bool {
	if owner == "self" && (s._backingNode.OwnerOfLeft != "" && s._backingNode.OwnerOfLeft != "self") {
		fmt.Println("\"", owner, "\" tried to set Left, but Node \"", s._backingNode.Id, "\" left is already owned by \"", s._backingNode.OwnerOfLeft, "\"")

		return true
	}

	return false
}

func hasTopOwner(s *ShadowNode, owner string) bool {
	if owner == "self" && (s._backingNode.OwnerOfTop != "" && s._backingNode.OwnerOfTop != "self") {
		fmt.Println("\"", owner, "\" tried to set Top, but Node \"", s._backingNode.Id, "\" top is already owned by \"", s._backingNode.OwnerOfTop, "\"")

		return true
	}

	return false
}

func (s *ShadowNode) DimWidth__Self_and_Node(width int, owner string) *ShadowNode {
	if !hasWidthOwner(s, owner) {
		s.Width = width

		s._backingNode.Width = s.Width
		s._backingNode.OwnerOfWidth = owner
	} else {
		fmt.Println("DimWidth__Self_and_Node ", )
	}

	return s
}

func (s *ShadowNode) DimWidth__Node_Only(owner string) *ShadowNode {
	if !hasWidthOwner(s, owner) {

		s._backingNode.Width = s.Width
		s._backingNode.OwnerOfWidth = owner
	} else {
		fmt.Println("DimWidth__Node_Only ", )
	}

	return s
}

func (s *ShadowNode) DimHeight__Self_and_Node(height int, owner string) *ShadowNode {
	if !hasHeightOwner(s, owner) {
		s.Height = height

		s._backingNode.Height = s.Height
		s._backingNode.OwnerOfHeight = owner
	} else {
		fmt.Println("DimHeight__Self_and_Node ", )
	}

	return s
}

func (s *ShadowNode) DimHeight__Node_Only(owner string) *ShadowNode {
	if !hasHeightOwner(s, owner) {
		s._backingNode.Height = s.Height
		s._backingNode.OwnerOfHeight = owner
	} else {
		fmt.Println("DimHeight__Node_Only ", )
	}

	return s
}

func (s *ShadowNode) PosLeft__Self_and_Node(left int, owner string) *ShadowNode {
	if !hasLeftOwner(s, owner) {
		s.Left = left

		s._backingNode.Left = s.Left
		s._backingNode.OwnerOfLeft = owner
	} else {
		fmt.Println("PosLeft__Self_and_Node ", )
	}

	return s
}

func (s *ShadowNode) PosLeft__Node_Only(owner string) *ShadowNode {
	if !hasLeftOwner(s, owner) {

		s._backingNode.Left = s.Left
		s._backingNode.OwnerOfLeft = owner
	} else {
		fmt.Println("PosLeft__Node_Only ", )
	}

	return s
}

func (s *ShadowNode) PosTop__Self_and_Node(top int, owner string) *ShadowNode {
	if !hasTopOwner(s, owner) {
		s.Top = top

		s._backingNode.Top = s.Top
		s._backingNode.OwnerOfTop = owner
	} else {
		fmt.Println("PosTop__Self_and_Node ", )
	}

	return s
}

func (s *ShadowNode) PosTop__Node_Only(owner string) *ShadowNode {
	if !hasTopOwner(s, owner) {
		s._backingNode.Top = s.Top
		s._backingNode.OwnerOfTop = owner
	} else {
		fmt.Println("PosTop__Node_Only ", )
	}

	return s
}



