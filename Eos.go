package bl

import (
	"github.com/amortaza/go-hal"
	"fmt"
)

var g_canvas map[string] hal.Canvas

func init() {

	g_canvas = make(map[string] hal.Canvas)
}

func getCanvas(node *Node) hal.Canvas {

	canvas, ok := g_canvas[node.Id]

	if !ok {

		canvas = g_graphics.NewCanvas(node.Width, node.Height)

		g_canvas[node.Id] = canvas

		// since we have a blank canvas, we need to flag that we need a redraw
		node.Dirty = true

		return canvas
	}

	// note: node.Dirty does not dictate whether we need a new canvas or not.
	// but a new canvas necesarily means dirty = true since we need a repaint.
	// dirty means that the content of the canvas needs repainting

	if canvas.GetWidth() != node.Width || canvas.GetHeight() != node.Height {

		// todo: if canvas is shrinking, no need to get a brand new canvas, re-use existing one

		canvas.Free()

		canvas = g_graphics.NewCanvas(node.Width, node.Height)

		g_canvas[node.Id] = canvas

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

	debug("Rendering canvas...")

	canvas.Begin()
	{
		canvas.Clear(.3, .3, .3)

		if !node.CustomsShouldRendersAfterKids {
			renderCustom(node)
		}

		for kide := node.Kids.Front(); kide != nil; kide = kide.Next() {

			kid := kide.Value.(*Node)

			kidCanvas := renderCanvas(kid)

			kidCanvas.Paint(false, kid.Left , kid.Top, FourOnesFloat32)
		}

		if node.CustomsShouldRendersAfterKids {
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

func freeCanvases() {

	for nodeId, canvas := range g_canvas {

		fmt.Println("(-) Freeing canvas for " + nodeId)

		canvas.Free()
	}
}