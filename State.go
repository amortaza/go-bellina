package bl

import (
	"adt"
	"g4"
)

var Root_Node *Node
var Current_Node *Node

var Mouse_X, Mouse_Y int32

var g_nodeByID map[string] *Node
var g_nodeByID_Previous map[string] *Node
var g_nodeStack adt.Stack
var g_textureByPartialName map[string] *g4.Texture

var FourOnesFloat = []float32{1,1,1,1}
var FourOnesInt = []int32{1,1,1,1}
