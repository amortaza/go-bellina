package bl

import "github.com/amortaza/go-g4"

type ShadowNode struct {
	ID string

	Left, Top, Width, Height int32

	Red1, Green1, Blue1 float32
	Red2, Green2, Blue2 float32

	NodeOpacity []float32

	Flags uint32

	Label string
	LabelOpacity float32

	FontName string
	FontSize int32
	FontRed, FontGreen, FontBlue float32
	FontNudgeX, FontNudgeY int32

	BorderThickness []int32
	BorderRed, BorderGreen, BorderBlue float32
	BorderTopsCanvas bool

	// no need to free this - this is globally managed
	Texture *g4.Texture

	SeeThru bool
}

func NewShadowNode(node *Node) *ShadowNode {
	shadow := &ShadowNode{
		node.ID,
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
		make([]int32, 4),
		node.BorderRed,node.BorderGreen,node.BorderBlue,
		node.BorderTopsCanvas,
		node.Texture,
		node.SeeThru,
	}

	copy(shadow.NodeOpacity, node.NodeOpacity)
	copy(shadow.BorderThickness, node.BorderThickness)

	return shadow
}