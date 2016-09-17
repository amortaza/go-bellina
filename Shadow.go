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
