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
		if node.CustomRender_Hook != nil {
			node.CustomRender_Hook(node)
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
