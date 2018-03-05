package bl

import (
	"fmt"
	"github.com/amortaza/go-hal"
)

var FourOnesFloat32 = []float32{1,1,1,1}

func DivId(id string) {
	Current_Node = GetNodeById(id)

	g_nodeStack.Push(Current_Node)
}

func GetNodeById(id string ) *Node {
	node, ok := g_nodeById[id]

	if !ok {
		fmt.Println("Unable to find Node by Id ", id )

		panic("See print out - GetNodeById error")
	}

	return node
}

func Stabilize(node *Node) {

	for e := node.funcs_pre_kids.Front(); e != nil; e = e.Next() {
		cb := e.Value.(func())
		cb()
	}

	for k := node.Kids.Front(); k != nil; k = k.Next() {
		kid := k.Value.(*Node)
		Stabilize(kid)
	}

	for e := node.funcs_post_kids.Front(); e != nil; e = e.Next() {
		cb := e.Value.(func())
		cb()
	}
}

func Disp(node *Node) {
	fmt.Println("Node ", node.Id, "(", node.Dirty, node.Left, ", ", node.Top, ") (", node.Width, " x ", node.Height, ")")
}

func SetMouseCursor(cursor hal.MouseCursor) {
	Hal.SetMouseCursor(cursor)
}
