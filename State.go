package bl

import (
	"github.com/amortaza/go-adt"
	"github.com/amortaza/go-g4"
	"container/list"
)

var Root_Node *Node
var Current_Node *Node

var Mouse_X, Mouse_Y int32

var g_nodeByID map[string] *Node
var g_nodeByID_Previous map[string] *Node
var g_nodeStack adt.Stack
var g_textureByPartialName map[string] *g4.Texture

var g_pluginParamsNodeId_int32 map[string] (map[string] int32)
var g_pluginParamsNodeId_string map[string] (map[string] string)

var g_pluginTicks *list.List

var FourOnesFloat = []float32{1,1,1,1}
var FourOnesInt = []int32{1,1,1,1}
var FourTwosInt = []int32{2,2,2,2}
