package bl

import (
	"container/list"
	"github.com/amortaza/go-adt"
	"fmt"
)

var g_nodeById map[string] *Node
var g_nodeStack adt.Stack
var g_funcs list.List

type Node struct {

	Id                       string

	Left, Top, Width, Height int

	Parent                   *Node
	Kids                     *list.List

	OnMouseMoveCallbacks     *list.List
	OnMouseButtonCallbacks   *list.List

	CustomRender             func(node *Node)
	CustomRenderTopsKids     bool

	InvisibleToEvents        bool

	OwnerOfLeft, OwnerOfTop, OwnerOfWidth, OwnerOfHeight string

	SettledBoundary bool
	SettledKids bool
}

func NewNode() *Node {
	node := &Node{}

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

func (node *Node) OwnLeft(owner string) bool {
	if node.OwnerOfLeft == "" || node.OwnerOfLeft == "*" {
		node.OwnerOfLeft = owner

	} else if owner != "*" && node.OwnerOfLeft != owner {
		fmt.Println("Node \"", node.Id, "\" Left is already owned by \"", node.OwnerOfLeft, "\" and set to ", node.Left, ", it cannot be owned by \"", owner, "\"")
		return false
	}

	return true
}

func (node *Node) OwnTop(owner string) bool {
	if node.OwnerOfTop == "" || node.OwnerOfTop == "*" {
		node.OwnerOfTop = owner

	} else if owner != "*" && node.OwnerOfTop != owner {
		fmt.Println("Node \"", node.Id, "\" Top is already owned by \"", node.OwnerOfTop, "\" and set to ", node.Top, ", it cannot be owned by \"", owner, "\"")
		return false
	}

	return true
}

func (node *Node) OwnWidth(owner string) bool {
	if node.OwnerOfWidth == ""  || node.OwnerOfWidth == "*" {
		node.OwnerOfWidth = owner
		
	} else if owner != "*" && node.OwnerOfWidth != owner {
		fmt.Println("Node \"", node.Id, "\" Width is already owned by \"", node.OwnerOfWidth, "\" and set to ", node.Width, ", it cannot be owned by \"", owner, "\"")
		return false
	}

	return true
}

func (node *Node) OwnHeight(owner string) bool {
	if node.OwnerOfHeight == "" || node.OwnerOfHeight == "*" {
		node.OwnerOfHeight = owner

	} else if owner != "*" && node.OwnerOfHeight != owner {
		fmt.Println("Node \"", node.Id, "\" Height is already owned by \"", node.OwnerOfHeight, "\" and set to ", node.Height, ", it cannot be owned by \"", owner, "\"")
		return false
	}

	return true
}

func (node *Node) OwnPos(owner string) bool {
	return node.OwnLeft(owner) && node.OwnTop(owner)
}

func (node *Node) OwnDim(owner string) bool {
	return node.OwnWidth(owner) && node.OwnHeight(owner)
}

func (node *Node) Free() {
	// free shadow node here!
	// todo
}