package bl

type ShadowNode struct {

	Id string

	Left, Top, Width, Height int

	BackingNode *Node // this is the node that the shadow is backing!!
}

func newShadowNode(node *Node) *ShadowNode {

	shadow := &ShadowNode{
		node.Id,
		node.Left, node.Top,
		node.Width, node.Height,
		nil, // this will get overwritten anyways, do not set to node
	}

	return shadow
}

func (shadow *ShadowNode) SetWidth_on_Self_and_Node(width int, owner string) *ShadowNode {

	if shadow.BackingNode.OwnsWidth(owner) {

		shadow.Width = width
		shadow.BackingNode.Width = shadow.Width
	}

	return shadow
}

func (shadow *ShadowNode) SetWidth_on_Node_Only(owner string) *ShadowNode {

	if shadow.BackingNode.OwnsWidth(owner) {
		shadow.BackingNode.Width = shadow.Width
	}

	return shadow
}

func (shadow *ShadowNode) SetHeight_on_Self_and_Node(height int, owner string) *ShadowNode {

	if shadow.BackingNode.OwnsHeight(owner) {
		shadow.Height = height
		shadow.BackingNode.Height = shadow.Height
	}

	return shadow
}

func (shadow *ShadowNode) SetDim_on_Self_and_Node(width, height int, owner string) *ShadowNode {

	if shadow.BackingNode.OwnsDim(owner) {
		shadow.Width, shadow.Height = width, height
		shadow.BackingNode.Width, shadow.BackingNode.Height = width, height
	}

	return shadow
}

func (shadow *ShadowNode) SetHeight_on_Node_Only(owner string) *ShadowNode {

	if shadow.BackingNode.OwnsHeight(owner) {
		shadow.BackingNode.Height = shadow.Height
	}

	return shadow
}

func (shadow *ShadowNode) SetDim_on_Node_Only(owner string) *ShadowNode {

	if shadow.BackingNode.OwnsDim(owner) {
		shadow.BackingNode.Width = shadow.Width
		shadow.BackingNode.Height = shadow.Height
	}

	return shadow
}

func (shadow *ShadowNode) SetLeft_on_Self_and_Node(left int, owner string) *ShadowNode {

	if shadow.BackingNode.OwnsLeft(owner) {
		shadow.Left = left
		shadow.BackingNode.Left = shadow.Left
	}

	return shadow
}

func (shadow *ShadowNode) SetLeft_on_Node_Only(owner string) *ShadowNode {

	if shadow.BackingNode.OwnsLeft(owner) {
		shadow.BackingNode.Left = shadow.Left
	}

	return shadow
}

func (shadow *ShadowNode) SetTop_on_Self_and_Node(top int, owner string) *ShadowNode {

	if shadow.BackingNode.OwnsTop(owner) {
		shadow.Top = top
		shadow.BackingNode.Top = shadow.Top
	}

	return shadow
}

func (shadow *ShadowNode) SetPos_on_Self_and_Node(left, top int, owner string) *ShadowNode {

	if shadow.BackingNode.OwnsPos(owner) {
		shadow.Left, shadow.Top = left, top

		shadow.BackingNode.Left, shadow.BackingNode.Top = left, top
	}

	return shadow
}

func (shadow *ShadowNode) SetTop_on_Node_Only(owner string) *ShadowNode {

	if shadow.BackingNode.OwnsTop(owner) {
		shadow.BackingNode.Top = shadow.Top
	}

	return shadow
}

func (shadow *ShadowNode) SetPos_on_Node_Only(owner string) *ShadowNode {

	if shadow.BackingNode.OwnsPos(owner) {
		shadow.BackingNode.Left = shadow.Left
		shadow.BackingNode.Top = shadow.Top
	}

	return shadow
}



