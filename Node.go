package bl

import (
	"container/list"
	"github.com/amortaza/go-adt"
	"fmt"
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

	// if true, then mouse events do not fire on this node
	InvisibleToMouseEvents   bool

	OwnerOfLeft              string
	OwnerOfTop               string
	OwnerOfWidth             string
	OwnerOfHeight            string

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

func (node *Node) OwnsLeft(owner string) bool {

	if node.OwnerOfLeft == "" || node.OwnerOfLeft == "*" {
		node.OwnerOfLeft = owner

	} else if owner != "*" && node.OwnerOfLeft != owner {
		fmt.Println("Node \"", node.Id, "\" left is already owned by \"", node.OwnerOfLeft, "\" and set to ", node.left, ", it cannot be owned by \"", owner, "\"")
		return false
	}

	return true
}

func (node *Node) OwnsTop(owner string) bool {

	if node.OwnerOfTop == "" || node.OwnerOfTop == "*" {
		node.OwnerOfTop = owner

	} else if owner != "*" && node.OwnerOfTop != owner {
		fmt.Println("Node \"", node.Id, "\" top is already owned by \"", node.OwnerOfTop, "\" and set to ", node.top, ", it cannot be owned by \"", owner, "\"")
		return false
	}

	return true
}

func (node *Node) OwnsWidth(owner string) bool {

	if node.OwnerOfWidth == ""  || node.OwnerOfWidth == "*" {
		node.OwnerOfWidth = owner
		
	} else if owner != "*" && node.OwnerOfWidth != owner {
		fmt.Println("Node \"", node.Id, "\" Width is already owned by \"", node.OwnerOfWidth, "\" and set to ", node.width, ", it cannot be owned by \"", owner, "\"")
		return false
	}

	return true
}

func (node *Node) OwnsHeight(owner string) bool {

	if node.OwnerOfHeight == "" || node.OwnerOfHeight == "*" {
		node.OwnerOfHeight = owner

	} else if owner != "*" && node.OwnerOfHeight != owner {
		fmt.Println("Node \"", node.Id, "\" height is already owned by \"", node.OwnerOfHeight, "\" and set to ", node.height, ", it cannot be owned by \"", owner, "\"")
		return false
	}

	return true
}

func (node *Node) OwnsPos(owner string) bool {

	left := node.OwnsLeft(owner)
	top := node.OwnsTop(owner)

	return  left && top
}

func (node *Node) OwnsDim(owner string) bool {

	width := node.OwnsWidth(owner)
	height := node.OwnsHeight(owner)

	return width && height
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