package bl

import (
	"fmt"
	"github.com/amortaza/go-hal"
)

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

func Disp(node *Node) {
	fmt.Println("Node ", node.Id, "(", node.Dirty, node.left, ", ", node.top, ") (", node.width, " x ", node.height, ")")
}

func SetMouseCursor(cursor hal.MouseCursor) {
	Hal.SetMouseCursor(cursor)
}

func GetNodeAbsolutePos(node *Node)(absX, absY int) {
	if node == nil {
		return 0, 0
	}

	absX, absY = node.left, node.top

	node = node.Parent

	for node != nil {
		absX += node.left;
		absY += node.top;
		node = node.Parent
	}

	return absX, absY
}

