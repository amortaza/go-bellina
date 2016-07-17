package bl

import (
	"github.com/amortaza/go-adt"
	"github.com/amortaza/go-g5"
	"container/list"
)

var Root_Node *Node
var Current_Node *Node

var Mouse_X, Mouse_Y int

var g_nodeByID map[string] *Node
var g_nodeByID_Previous map[string] *Node
var g_nodeStack adt.Stack
var g_textureByPartialName map[string] *g5.Texture

var g_pluginTicks *list.List

var FourOnesFloat = []float32{1,1,1,1}
var FourOnesInt = []int{1,1,1,1}
var FourTwosInt = []int{2,2,2,2}
var FourZeroesFloat = []float32{0,0,0,0}
