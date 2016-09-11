package bl

import "fmt"

type ShadowNode struct {
	Id                       string
	ParentId                 string

	Left, Top, Width, Height int

	BackingNode              *Node // this is the node that the shadow is backing!!
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
	if shadow.BackingNode.OwnWidth(owner) {
		shadow.Width = width
		shadow.BackingNode.Width = shadow.Width
	}

	return shadow
}

func (shadow *ShadowNode) Width__Node_Only(owner string) *ShadowNode {
	if shadow.BackingNode.OwnWidth(owner) {
		shadow.BackingNode.Width = shadow.Width
	}

	return shadow
}

func (shadow *ShadowNode) Height__Self_and_Node(height int, owner string) *ShadowNode {
	if shadow.BackingNode.OwnHeight(owner) {
		shadow.Height = height
		shadow.BackingNode.Height = shadow.Height
	}

	return shadow
}

func (shadow *ShadowNode) Dim__Self_and_Node(width, height int, owner string) *ShadowNode {
	if shadow.BackingNode.OwnDim(owner) {
		shadow.Width, shadow.Height = width, height
		shadow.BackingNode.Width, shadow.BackingNode.Height = width, height
	}

	return shadow
}

func (shadow *ShadowNode) Height__Node_Only(owner string) *ShadowNode {
	if shadow.BackingNode.OwnHeight(owner) {
		shadow.BackingNode.Height = shadow.Height
	}

	return shadow
}

func (shadow *ShadowNode) Dim__Node_Only(owner string) *ShadowNode {
	if shadow.BackingNode.OwnDim(owner) {
		shadow.BackingNode.Width = shadow.Width
		shadow.BackingNode.Height = shadow.Height
	}

	return shadow
}

func (shadow *ShadowNode) Left__Self_and_Node(left int, owner string) *ShadowNode {
	if shadow.BackingNode.OwnLeft(owner) {
		shadow.Left = left
		shadow.BackingNode.Left = shadow.Left
	}

	return shadow
}

func (shadow *ShadowNode) Left__Node_Only(owner string) *ShadowNode {
	if shadow.BackingNode.OwnLeft(owner) {
		shadow.BackingNode.Left = shadow.Left
	}

	return shadow
}

func (shadow *ShadowNode) Top__Self_and_Node(top int, owner string) *ShadowNode {
	if shadow.BackingNode.OwnTop(owner) {
		shadow.Top = top
		shadow.BackingNode.Top = shadow.Top
	}

	return shadow
}

func (shadow *ShadowNode) Pos__Self_and_Node(left, top int, owner string) *ShadowNode {
	if shadow.BackingNode.OwnPos(owner) {
		shadow.Left, shadow.Top = left, top

		shadow.BackingNode.Left, shadow.BackingNode.Top = left, top
	}

	return shadow
}

func (shadow *ShadowNode) Top__Node_Only(owner string) *ShadowNode {
	if shadow.BackingNode.OwnTop(owner) {
		shadow.BackingNode.Top = shadow.Top
	}

	return shadow
}

func (shadow *ShadowNode) Pos__Node_Only(owner string) *ShadowNode {
	if shadow.BackingNode.OwnPos(owner) {
		shadow.BackingNode.Left = shadow.Left
		shadow.BackingNode.Top = shadow.Top
	}

	return shadow
}



