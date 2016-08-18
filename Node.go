package bl

import (
	"container/list"
	"github.com/amortaza/go-adt"
)

var g_nodeByID map[string] *Node
var g_nodeStack adt.Stack

type Node struct {
	Id                       string

	Left, Top, Width, Height int

	Parent                   *Node
	Kids                     *list.List

	OnMouseMoveCallbacks     *list.List
	OnMouseButtonCallbacks   *list.List

	CustomRender             func(node *Node)

	PreventBubbling          bool
	InvisibleToEvents        bool
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

func (node *Node) Free() {
	// free shadow node here!
	// todo
}