package bl

import "container/list"

func Root() {
	Current_Node = NewNode()

	Current_Node.Id = "ROOT"

	g_nodeById["ROOT"] = Current_Node

	g_nodeStack.Push(Current_Node)

	Root_Node = Current_Node
}

func Div() {
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

func Pos(left, top int) {
	Current_Node.Left, Current_Node.Top = left, top
}

func Dim(width, height int) {
	Current_Node.Width, Current_Node.Height = width, height
}

func CustomRenderer(f func(node *Node), topsLabel bool) {
	Current_Node.CustomRender = f
}

func PreventBubbling() {
	Current_Node.PreventBubbling = true
}

func InvisibleToEvents() {
	Current_Node.InvisibleToEvents = true
}

func OnMouseMove(cb func(*MouseMoveEvent)) {
	onMouseMoveOnNode(Current_Node, cb)
}

func OnMouseButton(cb func(*MouseButtonEvent)) {
	onMouseButtonOnNode(Current_Node, cb)
}

func GetNodeById(id string ) *Node {
	node, _ := g_nodeById[id]

	return node
}

func onMouseMoveOnNode(node *Node, cb func(*MouseMoveEvent)) {
	if node.OnMouseMoveCallbacks == nil {
		node.OnMouseMoveCallbacks = list.New()
	}

	node.OnMouseMoveCallbacks.PushBack(cb);
}

func onMouseButtonOnNode(node *Node, cb func(*MouseButtonEvent)) {
	if node.OnMouseButtonCallbacks == nil {
		node.OnMouseButtonCallbacks = list.New()
	}

	node.OnMouseButtonCallbacks.PushBack(cb);
}

