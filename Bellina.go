package bl

func Root() {
	Current_Node = NewNode()

	nodeStack.Push(Current_Node)

	Root_Node = Current_Node
}

func Div() {
	parent := Current_Node

	Current_Node = NewNode()

	parent.Kids.PushBack(Current_Node)
	Current_Node.Parent = parent

	nodeStack.Push(Current_Node)
}

func End() {
	Current_Node = nodeStack.Pop().(*Node)
}

func Pos(left, top int32) {
	Current_Node.Left, Current_Node.Top = left, top
}

func Dim(width, height int32) {
	Current_Node.Width, Current_Node.Height = width, height
}

func Color(red,green,blue float32) {
	Current_Node.Red1, Current_Node.Green1, Current_Node.Blue1 = red, green, blue
}

func Color2(red,green,blue float32) {
	Current_Node.Red2, Current_Node.Green2, Current_Node.Blue2 = red, green, blue
}

func Flag(flag uint32) {
	Current_Node.Flags = flag
}

func Label(label string) {
	Current_Node.Label = label
}

func LabelOpacity(opacity float32) {
	Current_Node.LabelOpacity = opacity
}

func Font(fontName string, fontSize int32) {
	Current_Node.FontName, Current_Node.FontSize = fontName, fontSize
}

func FontColor(red, green, blue float32) {
	Current_Node.FontRed, Current_Node.FontGreen, Current_Node.FontBlue = red, green, blue
}

func FontNudge(x, y int32) {
	Current_Node.FontNudgeX, Current_Node.FontNudgeY = x, y
}

func BorderThickness(thickness []int32) {
	Current_Node.BorderThickness = thickness
}

func BorderColor(red, green, blue float32) {
	Current_Node.BorderRed, Current_Node.BorderGreen, Current_Node.BorderBlue = red, green, blue
}

func BorderTopsCanvas() {
	Current_Node.BorderTopsCanvas = true
}

func NodeOpacity(opacity []float32) {
	Current_Node.NodeOpacity = opacity
}




















































