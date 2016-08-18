package bl

func GetNodeAt__VisibleToEvents(x, y int) *Node {

	if Root_Node == nil {
		return nil
	}

	if contains(Root_Node, x, y) {
		return getNodeAt__visibleToEvents(Root_Node, x, y)
	}

	return nil
}

func GetNodeAbsolutePos(node *Node)(absX, absY int) {
	if node == nil {
		return 0, 0
	}

	absX, absY = node.Left, node.Top

	node = node.Parent

	for node != nil {
		absX += node.Left;
		absY += node.Top;
		node = node.Parent
	}

	return absX, absY
}

func contains(node *Node, x, y int) (ans bool) {
	ans = node.Left <= x && node.Top <= y && node.Left + node.Width > x && node.Top + node.Height > y

	return ans
}

func getNodeAt__visibleToEvents(root *Node, x, y int) *Node {
	x -= root.Left
	y -= root.Top

	for e := root.Kids.Back(); e != nil; e = e.Prev() {
		kid := e.Value.(*Node)

		if kid.InvisibleToEvents {
			continue
		}

		if contains(kid, x, y) {
			return getNodeAt__visibleToEvents(kid, x, y)
		}
	}

	return root
}

