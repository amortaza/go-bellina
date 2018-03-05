package bl

import "github.com/amortaza/go-hal"

var g_canvas map[string] hal.Canvas

func init() {

	g_canvas = make(map[string] hal.Canvas)
}

func getCanvas(node *Node) hal.Canvas {

	canvas, ok := g_canvas[node.Id]

	if !ok {

		canvas = g_graphics.NewCanvas(node.Width, node.Height)

		g_canvas[node.Id] = canvas

	} else if node.Dirty {

		if canvas.GetWidth() != node.Width || canvas.GetHeight() != node.Height {

			canvas.Free()

			canvas = g_graphics.NewCanvas(node.Width, node.Height)

			g_canvas[node.Id] = canvas
		}
	}

	return canvas
}

func renderCanvas(node *Node) hal.Canvas {

	canvas := getCanvas(node)

	if (!node.Dirty) {
		return canvas
	}

	canvas.Begin()
	{
		canvas.Clear(.3, .3, .3)

		if !node.CustomRenderTopsKids {
			renderCustom(node)
		}

		for kide := node.Kids.Front(); kide != nil; kide = kide.Next() {

			kid := kide.Value.(*Node)

			kidCanvas := renderCanvas(kid)

			kidCanvas.Paint(false, kid.Left , kid.Top, FourOnesFloat32)
		}

		if node.CustomRenderTopsKids {
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