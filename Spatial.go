package bl

func getNodeAt_VisibleToMouseEvents(x, y int) *Node {
	if Root_Node == nil {
		return nil
	}

	if _node_contains_point(Root_Node, x, y) {
		return _getNodeAt_VisibleToMouseEvents(Root_Node, x, y)
	}

	return nil
}


func _getNodeAt_VisibleToMouseEvents(root *Node, x, y int) *Node {
	x -= root.left
	y -= root.top

	for e := root.Kids.Back(); e != nil; e = e.Prev() {
		kid := e.Value.(*Node)

		if kid.Invisible_to_Mouse_Events {
			continue
		}

		if _node_contains_point(kid, x, y) {
			return _getNodeAt_VisibleToMouseEvents(kid, x, y)
		}
	}

	return root
}

func _node_contains_point(node *Node, x, y int) (ans bool) {
	ans = node.left <= x && node.top <= y && node.left+ node.width > x && node.top+ node.height > y

	return ans
}
