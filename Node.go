package bl

type Node struct {
	Left, Top, Width, Height int32

	Red1, Green1, Blue1 float32
	Red2, Green2, Blue2 float32

	Flags uint32

	Label string
	LabelOpacity float32

	FontName string
	FontSize int32
	FontRed, FontGreen, FontBlue float32
	FontNudgeX, FontNudgeY int32
}

func NewNode() *Node {
	node := &Node{}

	node.Flags = COLOR_SOLID
	node.LabelOpacity = 1

	return node
}
