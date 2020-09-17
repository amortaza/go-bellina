package bl

var g_lastFrame_nodeById map[string] *Node

// note that this function returns a node, this is an optimization, the last frame node is being looked up anyway so why not return it
func mark_new_and_changed_nodes_and_their_parents_dirty(node *Node) *Node {
	last_frame_node, node_existed_in_last_frame := g_lastFrame_nodeById[node.Id]

	if !node_existed_in_last_frame {

		// making parents dirty is done in child loop below
		node.Dirty = true

	} else if last_frame_node.width != node.width || last_frame_node.height != node.height {

		// making parents dirty is done in child loop below
		node.Dirty = true
	}

	for e := node.Kids.Front(); e != nil; e = e.Next() {
		kid := e.Value.(*Node)

		// this is an optimization, the last frame node is being looked up anyway so why not return it
		last_frame_kid := mark_new_and_changed_nodes_and_their_parents_dirty(kid)

		if kid.Dirty {
			// it is not enough to just set node.Dirty to true, because what if the grand-kid is dirty!
			mark_node_and_parents_dirty(node)

			// there is no point in further analysis to send node.Dirty to true, since it has been set here
			continue
		}

		if last_frame_kid != nil && (last_frame_kid.left != kid.left || last_frame_kid.top != kid.top) {
			mark_node_and_parents_dirty(node)
		}
	}

	// this is an optimization, the last frame node is being looked up anyway so why not return it
	return last_frame_node
}

func mark_node_and_parents_dirty(node *Node) {
	node.Dirty = true

	var p = node.Parent

	for true {
		if p == nil {
			return
		}

		if p.Dirty {
			return
		}

		p.Dirty = true

		p = p.Parent
	}
}

func garbage_collect_removed_nodes_and_mark_their_parents_dirty() {

	// find deleted nodes
	for nodeId, last_frame_node := range g_lastFrame_nodeById {
		_, node_exists_in_current_frame := g_nodeById[nodeId]

		if node_exists_in_current_frame {
			continue
		}

		// last_frame_node has been deleted since it exists in last frame, but not this one

		if last_frame_node.Parent != nil {

			// Parent is still from LAST FRAME
			// everything so far is from 'previous' frame, get into 'this' frame
			current_frame_parent, parent_exists_in_current_frame := g_nodeById[last_frame_node.Parent.Id]

			if parent_exists_in_current_frame {
				mark_node_and_parents_dirty(current_frame_parent)
			}
		}

		freeNode(last_frame_node)
	}
}
