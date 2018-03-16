package bl

var g_shadowNodeById map[string] *ShadowNode

func init() {

	g_shadowNodeById = make(map[string] *ShadowNode)
}

func EnsureShadow() *ShadowNode {

	shadow, ok := g_shadowNodeById[Current_Node.Id]

	if !ok {

		shadow = newShadowNode(Current_Node)

		g_shadowNodeById[Current_Node.Id] = shadow
	}

	shadow.BackingNode = Current_Node

	return shadow
}

func EnsureShadowById(id string) *ShadowNode {

	shadow, ok := g_shadowNodeById[id]

	node := GetNodeById(id)

	if !ok {

		shadow = newShadowNode(node)

		g_shadowNodeById[id] = shadow
	}

	shadow.BackingNode = node

	return shadow
}

func EnsureShadowByNode(node *Node) *ShadowNode {

	shadow, ok := g_shadowNodeById[node.Id]

	if !ok {

		shadow = newShadowNode(node)

		g_shadowNodeById[node.Id] = shadow
	}

	shadow.BackingNode = node

	return shadow
}

func HasShadowById(id string) (*ShadowNode, bool) {

	shadow, ok := g_shadowNodeById[id]

	return shadow, ok
}

func syncFromShadow(node *Node) {

	validateNodeId(node)

	shadow, ok := g_shadowNodeById[node.Id]

	if ok {

		shadow.BackingNode.left = shadow.Left
		shadow.BackingNode.top = shadow.Top
		shadow.BackingNode.width = shadow.Width
		shadow.BackingNode.height = shadow.Height
	}

	for e := node.Kids.Front(); e != nil; e = e.Next() {

		kid := e.Value.(*Node)

		syncFromShadow(kid)
	}
}

func validateNodeId(node *Node) {

	if node.Id == "" {

		parentId := "No Parent"

		if node.Parent != nil {
			parentId = node.Parent.Id
		}

		panic( "Node MUST have ID, parent was " + parentId )
	}
}
