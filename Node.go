package bl

import (
	"container/list"
	"github.com/amortaza/go-adt"
	"fmt"
)

var g_nodeById map[string] *Node
var g_nodeStack adt.Stack

type Node struct {

	Id                       string

	Left, Top, Width, Height int

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
		fmt.Println("Node \"", node.Id, "\" Left is already owned by \"", node.OwnerOfLeft, "\" and set to ", node.Left, ", it cannot be owned by \"", owner, "\"")
		return false
	}

	return true
}

func (node *Node) OwnsTop(owner string) bool {

	if node.OwnerOfTop == "" || node.OwnerOfTop == "*" {
		node.OwnerOfTop = owner

	} else if owner != "*" && node.OwnerOfTop != owner {
		fmt.Println("Node \"", node.Id, "\" Top is already owned by \"", node.OwnerOfTop, "\" and set to ", node.Top, ", it cannot be owned by \"", owner, "\"")
		return false
	}

	return true
}

func (node *Node) OwnsWidth(owner string) bool {

	if node.OwnerOfWidth == ""  || node.OwnerOfWidth == "*" {
		node.OwnerOfWidth = owner
		
	} else if owner != "*" && node.OwnerOfWidth != owner {
		fmt.Println("Node \"", node.Id, "\" Width is already owned by \"", node.OwnerOfWidth, "\" and set to ", node.Width, ", it cannot be owned by \"", owner, "\"")
		return false
	}

	return true
}

func (node *Node) OwnsHeight(owner string) bool {

	if node.OwnerOfHeight == "" || node.OwnerOfHeight == "*" {
		node.OwnerOfHeight = owner

	} else if owner != "*" && node.OwnerOfHeight != owner {
		fmt.Println("Node \"", node.Id, "\" Height is already owned by \"", node.OwnerOfHeight, "\" and set to ", node.Height, ", it cannot be owned by \"", owner, "\"")
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
