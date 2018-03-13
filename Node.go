package bl

import (
	"container/list"
	"github.com/amortaza/go-adt"
)

var g_nodeById map[string] *Node
var g_nodeStack adt.Stack

var g_nodes_are_immutable bool

type Node struct {

	Id                       string

	left, top, width, height int

	Parent                   *Node
	Kids                     *list.List

	OnMouseMoveCallbacks     *list.List
	OnMouseButtonCallbacks   *list.List

	CustomRender_1                func(node *Node)
	CustomRender_2                func(node *Node)
	CustomsShouldRendersAfterKids bool

	OwnerOfLeft              string
	OwnerOfTop               string
	OwnerOfWidth             string
	OwnerOfHeight            string

	// if true, then mouse events do not fire on this node
	InvisibleToMouseEvents   bool

	SettledBoundary          bool
	SettledKids              bool

	Dirty                    bool

	stabilize_funcs_pre_kids  *list.List
}

func newNode() *Node {

	node := &Node{}

	node.Id = adt.NewUUID()

	node.Kids = list.New()

	return node
}

func (node *Node) CallMouseMoveCallbacks(e *MouseMoveEvent) {

	if node.OnMouseMoveCallbacks != nil {

		for element := node.OnMouseMoveCallbacks.Front(); element != nil; element = element.Next() {

			cb := element.Value.(func(*MouseMoveEvent))
			cb(e)
		}
	}
}

func (node *Node) CallMouseButtonCallbacks(e *MouseButtonEvent) {

	if node.OnMouseButtonCallbacks != nil {

		for element := node.OnMouseButtonCallbacks.Front(); element != nil; element = element.Next() {

			cb := element.Value.(func(*MouseButtonEvent))

			cb(e)
		}
	}
}

func (node *Node) SetLeft(left int) {

	if g_nodes_are_immutable {
		panic("You are trying to Set the Left of something when we are in immutable mode!")
	}

	node.left = left
}

func (node *Node) Left() int {

	return node.left
}

func (node *Node) SetTop(top int) {

	if g_nodes_are_immutable {
		panic("You are trying to Set the Top of something when we are in immutable mode!")
	}

	node.top = top
}

func (node *Node) Top() int {

	return node.top
}

func (node *Node) SetWidth(width int) {

	if g_nodes_are_immutable {
		panic("You are trying to Set the Width of something when we are in immutable mode!")
	}

	node.width = width
}

func (node *Node) Width() int {

	return node.width
}

func (node *Node) SetHeight(height int) {

	if g_nodes_are_immutable {
		panic("You are trying to Set the Height of something when we are in immutable mode!")
	}

	node.height = height
}

func (node *Node) Height() int {

	return node.height
}

// Left

func (node *Node) IsOwnerOfLeft(owner string) bool {

	return node.OwnerOfLeft == "" || node.OwnerOfLeft == owner
}

func (node *Node) SetOwnerOfLeft(newOwner string) bool {

	if node.OwnerOfLeft == "" {
		node.OwnerOfLeft = newOwner

	} else if node.OwnerOfLeft != newOwner {
		debug("Node \"" + node.Id + "\" Left is already owned by \"" + node.OwnerOfLeft + "\" it cannot be owned by \"" + newOwner + "\"", "")

		return false
	}

	return true
}

// Top

func (node *Node) IsOwnerOfTop(owner string) bool {

	return node.OwnerOfTop == "" || node.OwnerOfTop == owner
}

func (node *Node) SetOwnerOfTop(newOwner string) bool {

	if node.OwnerOfTop == "" {
		node.OwnerOfTop = newOwner

	} else if node.OwnerOfTop != newOwner {
		debug("Node \"" + node.Id + "\" Top is already owned by \"" + node.OwnerOfTop + "\" it cannot be owned by \"" + newOwner + "\"", "")

		return false
	}

	return true
}

// Width

func (node *Node) IsOwnerOfWidth(owner string) bool {

	return node.OwnerOfWidth == "" || node.OwnerOfWidth == owner
}

func (node *Node) SetOwnerOfWidth(newOwner string) bool {

	if node.OwnerOfWidth == "" {
		node.OwnerOfWidth = newOwner

	} else if node.OwnerOfWidth != newOwner {
		debug("Node \"" + node.Id + "\" Width is already owned by \"" + node.OwnerOfWidth + "\" it cannot be owned by \"" + newOwner + "\"", "")

		return false
	}

	return true
}

// Height

func (node *Node) IsOwnerOfHeight(owner string) bool {

	return node.OwnerOfHeight == "" || node.OwnerOfHeight == owner
}

func (node *Node) SetOwnerOfHeight(newOwner string) bool {

	if node.OwnerOfHeight == "" {
		node.OwnerOfHeight = newOwner

	} else if node.OwnerOfHeight != newOwner {
		debug("Node \"" + node.Id + "\" Height is already owned by \"" + node.OwnerOfHeight + "\" it cannot be owned by \"" + newOwner + "\"", "")

		return false
	}

	return true
}

