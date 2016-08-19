package bl

var g_shadowNodeById map[string] *ShadowNode

func init() {
	g_shadowNodeById = make(map[string] *ShadowNode)
}

// MUST return sing value cause fluent
func EnsureShadow() *ShadowNode {

	shadow, ok := g_shadowNodeById[Current_Node.Id]

	if !ok {

		shadow = NewShadowNode(Current_Node)

		g_shadowNodeById[Current_Node.Id] = shadow
	}

	shadow._backingNode = Current_Node

	return shadow
}

// MUST remain fluent - only one return value
func EnsureShadowById(id string) *ShadowNode {

	shadow, ok := g_shadowNodeById[id]

	node := GetNodeById(id)

	if !ok {

		shadow = NewShadowNode(node)

		g_shadowNodeById[id] = shadow
	}

	shadow._backingNode = node

	return shadow
}

// MUST remain fluent - only one return value
func EnsureShadowByNode(node *Node) *ShadowNode {

	shadow, ok := g_shadowNodeById[node.Id]

	if !ok {

		shadow = NewShadowNode(node)

		g_shadowNodeById[node.Id] = shadow
	}

	shadow._backingNode = node

	return shadow
}

func TestShadowById(id string) (*ShadowNode, bool) {
	shadow, ok := g_shadowNodeById[id]

	return shadow, ok
}
