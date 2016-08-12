package bl

import (
	"github.com/amortaza/go-g5"
)

func renderCanvas(node *Node) *g5.Canvas {

	if node == nil {
		return nil
	}

	canvas := g5.NewCanvas(node.Width, node.Height)

	canvas.Begin()
	{
		var color1 []float32

		color1 = []float32{.33, .33, .33, 1}

		if node.Id == "ROOT" {
			color1 = []float32{.1, 0, .1, 1}
		}

		if node.Id == "red" {
			//color1 = []float32{1, 0, 0, 1}
		}

		if node.Id == "green" {
			color1 = []float32{0, 1, 0, 1}
		}

		if node.Id == "blue" {
			color1 = []float32{0, 0, 1, 1}
		}

		//if node.CustomRender_Hook == nil {
			g5.DrawColorRect4v(0, 0, node.Width, node.Height, color1, color1, color1, color1)
		//}

		if node.CustomRender_Hook != nil {
			//ux.Ctx.BeginFrame(node.Width / 2, node.Height / 2, float32( 2 ))
			node.CustomRender_Hook(node)
			//ux.Ctx.EndFrame()
		}

		for kide := node.Kids.Front(); kide != nil; kide = kide.Next() {

			kid := kide.Value.(*Node)

			kidCanvas := renderCanvas(kid)

			kidCanvas.Paint(false, kid.Left , kid.Top, g5.FourOnesFloat32)

			kidCanvas.Free()
		}
	}
	canvas.End()

	return canvas
}
