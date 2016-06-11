package bl

import (
	"adt"
	"g4"
)

var Root_Node *Node
var Current_Node *Node

var Mouse_X, Mouse_Y int32

var g_nodeByID map[string] *Node
var g_nodeStack adt.Stack
var g_textureByPartialName map[string] *g4.Texture

var FourOnes = []float32{1,1,1,1}
