package bl

import (
	"github.com/amortaza/go-g4"
	"container/list"
	"github.com/amortaza/go-bellina/core"
)

func Root() {
	Current_Node = NewNode()

	Current_Node.Id = "ROOT"

	g_nodeByID["ROOT"] = Current_Node

	g_nodeStack.Push(Current_Node)

	Root_Node = Current_Node
}

func Div() {
	parent := Current_Node

	Current_Node = NewNode()

	parent.Kids.PushBack(Current_Node)
	Current_Node.Parent = parent

	g_nodeStack.Push(Current_Node)
}

func ID(id string) {
	Current_Node.Id = id
	g_nodeByID[id] = Current_Node
}

func End() {
	g_nodeStack.Pop()

	if g_nodeStack.Size == 0 {
		Current_Node = nil
	} else {
		Current_Node = g_nodeStack.Top().(*Node)
	}
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

func GetFont() (string, int32) {
	return Current_Node.FontName, Current_Node.FontSize
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

func NodeOpacity1f(opacity float32) {
	Current_Node.NodeOpacity = []float32{opacity,opacity,opacity,opacity}
}

func NodeOpacity4f(opacity []float32) {
	Current_Node.NodeOpacity = opacity
}

func Texture(partialname string) {
 	texture, ok := g_textureByPartialName[partialname]

	if !ok {
		texture = g4.NewTexture()
		texture.LoadImage("github.com/amortaza/go-bellina-examples/assets/images/" + partialname + ".png")

		g_textureByPartialName[partialname] = texture
	}

	Current_Node.Texture = texture

	Current_Node.Width = texture.Width
	Current_Node.Height = texture.Height

	Current_Node.SeeThru = true
}

func OnMouseMove(cb func(*MouseMoveEvent)) {
	OnMouseMoveOnNode(Current_Node, cb)
}

func OnMouseMoveOnNode(node *Node, cb func(*MouseMoveEvent)) {
	if node.OnMouseMoveCallbacks == nil {
		node.OnMouseMoveCallbacks = list.New()
	}

	node.OnMouseMoveCallbacks.PushBack(cb);
}

func OnMouseButton(cb func(*MouseButtonEvent)) {
	OnMouseButtonOnNode(Current_Node, cb)
}

func OnMouseButtonOnNode(node *Node, cb func(*MouseButtonEvent)) {
	if node.OnMouseButtonCallbacks == nil {
		node.OnMouseButtonCallbacks = list.New()
	}

	node.OnMouseButtonCallbacks.PushBack(cb);
}

func GetNodeById(id string ) *Node {
	node, _ := g_nodeByID[id]

	return node
}

func GetFontHeight() int32 {

	fontname, fontsize := Current_Node.FontName, Current_Node.FontSize

	g4font := core.GetG4Font(fontname, fontsize)

	return g4font.Height
}

func SetI(pluginName, param string, value int32) {
	paramByPlugin, ok := g_pluginParamsNodeId_int32[Current_Node.Id]

	if !ok {
		paramByPlugin = make(map[string] int32)

		g_pluginParamsNodeId_int32[Current_Node.Id] = paramByPlugin
	}

	paramByPlugin[ pluginName + ":" + param] = value
}

func GetI(pluginName, param string) int32 {
	return GetI_fromNodeId(Current_Node.Id, pluginName, param)
}

func GetI_fromNodeId(nodeId, pluginName, param string) int32 {
	paramByPlugin, ok := g_pluginParamsNodeId_int32[nodeId]

	if !ok {
		return 0
	}

	value, ok2 := paramByPlugin[ pluginName + ":" + param]

	if !ok2 {
		return 0
	}

	return value
}

func Rect1(left, top, width, height int32, red, green, blue, opacity float32, topsLabel bool) {
	Current_Node.Rect1_LTWH = []int32{left,top,width,height}
	Current_Node.Rect1_TopsLabel = topsLabel
	Current_Node.Rect1_RGBA = []float32{red, green, blue, opacity}
}