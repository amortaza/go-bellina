package bl

var g_shadowNodeByID map[string] *ShadowNode

func EnsureShadow() (*ShadowNode) {
	return EnsureShadowByID(Current_Node.Id)
}

func EnsureShadowByID(id string) (*ShadowNode) {

	shadow, ok := g_shadowNodeByID[id]

	if !ok {
		node := GetNodeByID(id)
		shadow = NewShadowNode(node)
		g_shadowNodeByID[id] = shadow
	}

	return shadow
}

func GetShadowById(nodeid string) (*ShadowNode, bool) {
	shadow, ok := g_shadowNodeByID[nodeid]

	return shadow, ok
}

func GetShadow() (*ShadowNode, bool) {
	return GetShadowById(Current_Node.Id)
}