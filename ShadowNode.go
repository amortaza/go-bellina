package bl

import "fmt"

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

func (shadow *ShadowNode) SetWidth_on_Self_and_Node(width int, owner string) *ShadowNode {

	if g_nodes_are_immutable {
		panic("Cannot write on immutable")
	}

	if shadow.BackingNode.IsOwnerOfWidth(owner) {

		shadow.Width = width
		shadow.BackingNode.width = shadow.Width
	} else {
		fmt.Println("Wrong owner")
	}

	return shadow
}

func (shadow *ShadowNode) SetWidth_on_Node_Only(owner string) *ShadowNode {

	if g_nodes_are_immutable {
		panic("Cannot write on immutable")
	}

	if shadow.BackingNode.IsOwnerOfWidth(owner) {
		shadow.BackingNode.width = shadow.Width
	} else {
		fmt.Println("Wrong owner")
	}

	return shadow
}

func (shadow *ShadowNode) SetHeight_on_Self_and_Node(height int, owner string) *ShadowNode {

	if g_nodes_are_immutable {
		panic("Cannot write on immutable")
	}

	if shadow.BackingNode.IsOwnerOfHeight(owner) {
		shadow.Height = height
		shadow.BackingNode.height = shadow.Height
	} else {
		fmt.Println("Wrong owner")
	}

	return shadow
}

func (shadow *ShadowNode) SetHeight_on_Node_Only(owner string) *ShadowNode {

	if g_nodes_are_immutable {
		panic("Cannot write on immutable")
	}

	if shadow.BackingNode.IsOwnerOfHeight(owner) {
		shadow.BackingNode.height = shadow.Height
	} else {
		fmt.Println("Wrong owner")
	}

	return shadow
}

func (shadow *ShadowNode) SetLeft_on_Self_and_Node(left int, owner string) *ShadowNode {

	if g_nodes_are_immutable {
		panic("Cannot write on immutable")
	}

	if shadow.BackingNode.IsOwnerOfLeft(owner) {
		shadow.Left = left
		shadow.BackingNode.left = shadow.Left
	} else {
		fmt.Println("Wrong owner")
	}

	return shadow
}

func (shadow *ShadowNode) SetLeft_on_Node_Only(owner string) *ShadowNode {

	if g_nodes_are_immutable {
		panic("Cannot write on immutable")
	}

	if shadow.BackingNode.IsOwnerOfLeft(owner) {
		shadow.BackingNode.left = shadow.Left
	} else {
		fmt.Println("Wrong owner")
	}

	return shadow
}

func (shadow *ShadowNode) SetTop_on_Self_and_Node(top int, owner string) *ShadowNode {

	if g_nodes_are_immutable {
		panic("Cannot write on immutable")
	}

	if shadow.BackingNode.IsOwnerOfTop(owner) {
		shadow.Top = top
		shadow.BackingNode.top = shadow.Top
	} else {
		fmt.Println("Wrong owner")
	}

	return shadow
}

func (shadow *ShadowNode) SetTop_on_Node_Only(owner string) *ShadowNode {

	if g_nodes_are_immutable {
		panic("Cannot write on immutable")
	}

	if shadow.BackingNode.IsOwnerOfTop(owner) {
		shadow.BackingNode.top = shadow.Top
	} else {
		fmt.Println("Wrong owner")
	}

	return shadow
}

