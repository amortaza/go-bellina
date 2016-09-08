package bl

import "fmt"

type ShadowNode struct {
	Id                       string
	ParentId                 string

	Left, Top, Width, Height int

	backingNode              *Node // this is the node that the shadow is backing!!
}

func fake4() {
	var _ = fmt.Print
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

func (shadow *ShadowNode) Width__Self_and_Node(width int, owner string) *ShadowNode {
	if owner == "*" || shadow.backingNode.OwnWidth(owner) {
		shadow.Width = width
		shadow.backingNode.Width = shadow.Width
	}

	return shadow
}

func (shadow *ShadowNode) Width__Node_Only(owner string) *ShadowNode {
	if owner == "*" || shadow.backingNode.OwnWidth(owner) {
		shadow.backingNode.Width = shadow.Width
	}

	return shadow
}

func (shadow *ShadowNode) Height__Self_and_Node(height int, owner string) *ShadowNode {
	if owner == "*" || shadow.backingNode.OwnHeight(owner) {
		shadow.Height = height
		shadow.backingNode.Height = shadow.Height
	}

	return shadow
}

func (shadow *ShadowNode) Dim__Self_and_Node(width, height int, owner string) *ShadowNode {
	if owner == "*" || shadow.backingNode.OwnDim(owner) {
		shadow.Width, shadow.Height = width, height
		shadow.backingNode.Width, shadow.backingNode.Height = width, height
	}

	return shadow
}

func (shadow *ShadowNode) Height__Node_Only(owner string) *ShadowNode {
	if owner == "*" || shadow.backingNode.OwnHeight(owner) {
		shadow.backingNode.Height = shadow.Height
	}

	return shadow
}

func (shadow *ShadowNode) Dim__Node_Only(owner string) *ShadowNode {
	if owner == "*" || shadow.backingNode.OwnDim(owner) {
		shadow.backingNode.Width = shadow.Width
		shadow.backingNode.Height = shadow.Height
	}

	return shadow
}

func (shadow *ShadowNode) Left__Self_and_Node(left int, owner string) *ShadowNode {
	if owner == "*" || shadow.backingNode.OwnLeft(owner) {
		shadow.Left = left
		shadow.backingNode.Left = shadow.Left
	}

	return shadow
}

func (shadow *ShadowNode) Left__Node_Only(owner string) *ShadowNode {
	if owner == "*" || shadow.backingNode.OwnLeft(owner) {
		shadow.backingNode.Left = shadow.Left
	}

	return shadow
}

func (shadow *ShadowNode) Top__Self_and_Node(top int, owner string) *ShadowNode {
	if owner == "*" || shadow.backingNode.OwnTop(owner) {
		shadow.Top = top
		shadow.backingNode.Top = shadow.Top
	}

	return shadow
}

func (shadow *ShadowNode) Pos__Self_and_Node(left, top int, owner string) *ShadowNode {
	if owner == "*" || shadow.backingNode.OwnPos(owner) {
		shadow.Left, shadow.Top = left, top

		shadow.backingNode.Left, shadow.backingNode.Top = left, top
	}

	return shadow
}

func (shadow *ShadowNode) Top__Node_Only(owner string) *ShadowNode {
	if owner == "*" || shadow.backingNode.OwnTop(owner) {
		shadow.backingNode.Top = shadow.Top
	}

	return shadow
}

func (shadow *ShadowNode) Pos__Node_Only(owner string) *ShadowNode {
	if owner == "*" || shadow.backingNode.OwnPos(owner) {
		shadow.backingNode.Left = shadow.Left
		shadow.backingNode.Top = shadow.Top
	}

	return shadow
}



