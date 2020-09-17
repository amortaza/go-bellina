package bl

import (
	"container/list"
	"github.com/amortaza/go-adt"
	"github.com/amortaza/go-bellina/funclist"
)

var g_nodeById map[string] *Node
var g_nodeStack adt.Stack

var g_nodes_are_immutable bool

type Node struct {

	Id                       string

	left, top, width, height int

	Parent                   *Node
	Kids                     *list.List

	On_Mouse_Move_Callbacks   *list.List
	On_Mouse_Button_Callbacks *list.List

	// Use case for 1/2 custom renderer is the Border.Fill + Border.Wire
	CustomRender_1            func(node *Node)
	CustomRender_2            func(node *Node)
	Customs_Render_After_Kids bool

	// if true, then mouse events do not fire on this node
	Invisible_to_Mouse_Events bool

	//SettledBoundary          bool
	//SettledKids              bool

	Dirty                     bool

	stabilize_funcs_pre_kids  *funclist.FuncList
	stabilize_funcs_post_kids *funclist.FuncList
}

func newNode() *Node {

	node := &Node{}

	node.Kids = list.New()

	return node
}

func (node *Node) CallMouseMoveCallbacks(e *MouseMoveEvent) {

	if node.On_Mouse_Move_Callbacks != nil {

		for element := node.On_Mouse_Move_Callbacks.Front(); element != nil; element = element.Next() {

			cb := element.Value.(func(*MouseMoveEvent))
			cb(e)
		}
	}
}

func (node *Node) CallMouseButtonCallbacks(e *MouseButtonEvent) {

	if node.On_Mouse_Button_Callbacks != nil {

		for element := node.On_Mouse_Button_Callbacks.Front(); element != nil; element = element.Next() {

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

