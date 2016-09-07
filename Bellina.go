package bl

import (
	"container/list"
	"fmt"
)

func fake() {
	var _ = fmt.Println
}

func Root() {
	g_root_depth++

	Current_Node = NewNode()

	Current_Node.Id = "ROOT"

	g_nodeById["ROOT"] = Current_Node

	g_nodeStack.Push(Current_Node)

	Root_Node = Current_Node
}

func Div() {
	g_root_depth++

	parent := Current_Node

	Current_Node = NewNode()

	parent.Kids.PushBack(Current_Node)
	Current_Node.Parent = parent

	g_nodeStack.Push(Current_Node)
}

func DivId(id string) {
	Current_Node = GetNodeById(id)

	g_nodeStack.Push(Current_Node)
}

func Id(id string) {
	Current_Node.Id = id
	g_nodeById[id] = Current_Node
}

func End() {
	g_nodeStack.Pop()

	if g_nodeStack.Size == 0 {
		Current_Node = nil
	} else {
		Current_Node = g_nodeStack.Top().(*Node)
	}
}

func OwnLeft(owner string) {
	if Current_Node.OwnerOfLeft == "" {
		Current_Node.OwnerOfLeft = owner
		
	} else {
		fmt.Println("Node \"", Current_Node.Id, "\" Left is already owned by \"", Current_Node.OwnerOfLeft, "\" and set to ", Current_Node.Left, ", it cannot be owned by ", owner)
	}
}

func OwnTop(owner string) {
	if Current_Node.OwnerOfTop == "" {
		Current_Node.OwnerOfTop = owner

	} else {
		fmt.Println("Node \"", Current_Node.Id, "\" Top is already owned by \"", Current_Node.OwnerOfTop, "\" and set to ", Current_Node.Top, ", it cannot be owned by ", owner)
	}
}

func OwnWidth(owner string) {
	if Current_Node.OwnerOfWidth == "" {
		Current_Node.OwnerOfWidth = owner
		
	} else {
		fmt.Println("Node \"", Current_Node.Id, "\" Width is already owned by \"", Current_Node.OwnerOfWidth, "\" and set to ", Current_Node.Width, ", it cannot be owned by ", owner)
	}
}

func OwnHeight(owner string) {
	if Current_Node.OwnerOfHeight == "" {
		Current_Node.OwnerOfHeight = owner

	} else {
		fmt.Println("Node \"", Current_Node.Id, "\" Height is already owned by \"", Current_Node.OwnerOfHeight, "\" and set to ", Current_Node.Height, ", it cannot be owned by ", owner)
	}
}

func Pos(left, top int) {

	Current_Node.Left = left
	Current_Node.Top = top
}

func Dim(width, height int) {
		Current_Node.Width = width
		Current_Node.Height = height
}

func CustomRenderer(f func(node *Node), topsKids bool) {
	Current_Node.CustomRender = f
	Current_Node.CustomRenderTopsKids = topsKids
}

func PreventBubbling() {
	Current_Node.PreventBubbling = true
}

func InvisibleToEvents() {
	Current_Node.InvisibleToEvents = true
}

func RequireSettledBoundaries()  {
	if !Current_Node.SettledBoundary {
		fmt.Println("Boundary has not been settled for node ", Current_Node.Id)
		panic("See print out - RequireSettledBoundaries error")
	}
}

func RequireSettledKids() {
	if !Current_Node.SettledKids {
		fmt.Println("Kids have not been settled for node ", Current_Node.Id)
		panic("See print out - RequireSettledKids error")
	}
}

func SettleBoundary() {
	Current_Node.SettledBoundary = true
}

func SettleKids() {
	Current_Node.SettledKids = true
}

func OnMouseMove(cb func(*MouseMoveEvent)) {
	OnMouseMoveOnNode(Current_Node, cb)
}

func OnMouseButton(cb func(*MouseButtonEvent)) {
	OnMouseButtonOnNode(Current_Node, cb)
}

func GetNodeById(id string ) *Node {
	node, ok := g_nodeById[id]

	if !ok {
		fmt.Println("Unable to find Node by Id ", id )

		panic("See print out - GetNodeById error")
	}

	return node
}

func OnMouseMoveOnNode(node *Node, cb func(*MouseMoveEvent)) {
	if node.OnMouseMoveCallbacks == nil {
		node.OnMouseMoveCallbacks = list.New()
	}

	node.OnMouseMoveCallbacks.PushBack(cb);
}

func OnMouseButtonOnNode(node *Node, cb func(*MouseButtonEvent)) {
	if node.OnMouseButtonCallbacks == nil {
		node.OnMouseButtonCallbacks = list.New()
	}

	node.OnMouseButtonCallbacks.PushBack(cb);
}

func AddFunc(f func()) {
	g_funcs.PushBack(f)
}

func Disp(n *Node) {
	fmt.Println("Node ", n.Id, "(", n.Left, ", ", n.Top, ") (", n.Width, " x ", n.Height, ")")
}
