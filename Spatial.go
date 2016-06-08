package bl

func contains(node *Node, x, y int32) (ans bool) {
    	ans = node.Left <= x && node.Top <= y && node.Left + node.Width > x && node.Top + node.Height > y

	return ans
}

func getNodeAt(root *Node, x, y int32) *Node {
	x -= root.Left
	y -= root.Top

	for e := root.Kids.Front(); e != nil; e = e.Next() {
		kid := e.Value.(*Node)

		if contains(kid, x, y) {
			return getNodeAt(kid, x, y)
		}
	}

	return root
}

func GetNodeAt(x, y int32) *Node {

	if contains(Root_Node, x, y) {
		return getNodeAt(Root_Node, x, y)
	}

	return nil
}
