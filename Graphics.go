package bl

import (
	"github.com/amortaza/go-bellina/debug"
	"github.com/amortaza/go-hal"
)

var g_canvasById map[string] hal.Canvas
var g_four_ones = []float32{1,1,1,1}

func init() {
	g_canvasById = make(map[string] hal.Canvas)
}

func getCanvas(node *Node) hal.Canvas {
	canvas, ok := g_canvasById[node.Id]

	if !ok {
		debug.Log("  Allocating Canvas for " + node.Id, debug.Canvas)

		canvas = g_graphics.NewCanvas(node.width, node.height)

		g_canvasById[node.Id] = canvas

		// since we have a blank canvas, we need to flag that we need a redraw
		node.Dirty = true

		return canvas
	}

	// note: node.Dirty does not dictate whether we need a new canvas or not.
	// but a new canvas necesarily means dirty = true since we need a repaint.
	// dirty means that the content of the canvas needs repainting

	if canvas.GetWidth() != node.width || canvas.GetHeight() != node.height {

		// todo: if canvas is shrinking, no need to get a brand new canvas, re-use existing one

		canvas.Free()

		debug.Log("  Allocating NEW for OLD Canvas for " + node.Id, debug.Canvas)

		canvas = g_graphics.NewCanvas(node.width, node.height)

		g_canvasById[node.Id] = canvas

		// since we have a blank canvas, we need to flag that we need a redraw
		node.Dirty = true
	}

	return canvas
}

func renderCanvas(node *Node) hal.Canvas {
	canvas := getCanvas(node)

	if (!node.Dirty) {
		return canvas
	}

	debug.Log("Rendering " + node.Id, debug.Canvas )

	canvas.Begin()
	{
		canvas.Clear(.15, .3, .3)

		if !node.Customs_Render_After_Kids {
			renderCustom(node)
		}

		for kide := node.Kids.Front(); kide != nil; kide = kide.Next() {
			kid := kide.Value.(*Node)

			kidCanvas := renderCanvas(kid)

			kidCanvas.Paint(true, kid.left, kid.top, g_four_ones)
		}

		if node.Customs_Render_After_Kids {
			renderCustom(node)
		}
	}

	canvas.End()

	return canvas
}

func renderCustom(node *Node) {
	if node.CustomRender_1 != nil {
		node.CustomRender_1(node)
	}

	if node.CustomRender_2 != nil {
		node.CustomRender_2(node)
	}
}

