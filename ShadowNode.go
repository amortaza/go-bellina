package bl

import "github.com/amortaza/go-g4"

type ShadowNode struct {
	Id                                 string
	ParentID                           string

	Left, Top, Width, Height           int

	Red1, Green1, Blue1                float32
	Red2, Green2, Blue2                float32

	NodeOpacity                        []float32

	Flags                              uint32

	Label                              string
	LabelOpacity                       float32

	FontName                           string
	FontSize                           int
	FontRed, FontGreen, FontBlue       float32
	FontNudgeX, FontNudgeY             int

	BorderThickness                    []int
	BorderRed, BorderGreen, BorderBlue float32
	BorderTopsCanvas                   bool

	// no need to free this - this is globally managed
	Texture                            *g4.Texture

	SeeThru                            bool

	_backingNode *Node // only valud during calls to Color1()
}

func NewShadowNode(node *Node) *ShadowNode {
	var parentId string

	if node.Parent != nil {
		parentId = node.Parent.Id
	}

	shadow := &ShadowNode{
		node.Id,
		parentId,
		node.Left, node.Top, node.Width, node.Height,
		node.Red1, node.Green1,node.Blue1,
		node.Red2, node.Green2,node.Blue2,
		make([]float32, 4),
		node.Flags,
		node.Label,
		node.LabelOpacity,
		node.FontName,
		node.FontSize,
		node.FontRed,node.FontGreen,node.FontBlue,
		node.FontNudgeX, node.FontNudgeY,
		make([]int, 4),
		node.BorderRed,node.BorderGreen,node.BorderBlue,
		node.BorderTopsCanvas,
		node.Texture,
		node.SeeThru,
		nil,
	}

	copy(shadow.NodeOpacity, node.NodeOpacity)
	copy(shadow.BorderThickness, node.BorderThickness)

	return shadow
}