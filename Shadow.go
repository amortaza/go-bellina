package bl

var g_shadowNodeByID map[string] *ShadowNode

func EnsureShadow() (*ShadowNode) {

	shadow, ok := g_shadowNodeByID[Current_Node.ID]

	if !ok {
		shadow = NewShadowNode(Current_Node)
		g_shadowNodeByID[Current_Node.ID] = shadow
	}

	return shadow
}

func GetShadow(nodeid string) (*ShadowNode, bool) {
	shadow, ok := g_shadowNodeByID[nodeid]

	return shadow, ok
}