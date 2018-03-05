package bl

func getNodeAt__VisibleToMouseEvents(x, y int) *Node {

	if Root_Node == nil {
		return nil
	}

	if _contains(Root_Node, x, y) {
		return _getNodeAt_VisibleToMouseEvents(Root_Node, x, y)
	}

	return nil
}


func _getNodeAt_VisibleToMouseEvents(root *Node, x, y int) *Node {

	x -= root.Left
	y -= root.Top

	for e := root.Kids.Back(); e != nil; e = e.Prev() {

		kid := e.Value.(*Node)

		if kid.InvisibleToMouseEvents {
			continue
		}

		if _contains(kid, x, y) {
			return _getNodeAt_VisibleToMouseEvents(kid, x, y)
		}
	}

	return root
}

func _contains(node *Node, x, y int) (ans bool) {

	ans = node.Left <= x && node.Top <= y && node.Left + node.Width > x && node.Top + node.Height > y

	return ans
}
