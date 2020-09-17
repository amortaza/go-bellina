package bl

type spacial struct {}

var g_spacial = &spacial{}

func (s *spacial) GetNodeAt_VisibleToMouseEvents(x, y int) *Node {
	if Root_Node == nil {
		return nil
	}

	if s.node_contains_point(Root_Node, x, y) {
		return s.getNodeAt_VisibleToMouseEvents(Root_Node, x, y)
	}

	return nil
}

func (s *spacial) getNodeAt_VisibleToMouseEvents(root *Node, x, y int) *Node {
	x -= root.left
	y -= root.top

	// note that we are moving in reverse order - Painters Algorithm
	for e := root.Kids.Back(); e != nil; e = e.Prev() {
		kid := e.Value.(*Node)

		if kid.Invisible_to_Mouse_Events {
			continue
		}

		if s.node_contains_point(kid, x, y) {
			return s.getNodeAt_VisibleToMouseEvents(kid, x, y)
		}
	}

	return root
}

func (s *spacial) node_contains_point(node *Node, x, y int) (ans bool) {
	ans = node.left <= x &&
		  node.top <= y &&
		  node.left + node.width > x &&
		  node.top + node.height > y

	return ans
}
