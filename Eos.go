package bl

import (
	"github.com/amortaza/go-g5"
	"fmt"
)

func fake2() {
    var _ = fmt.Println
}

func renderCanvas(node *Node) *g5.Canvas {

	if node == nil {
		return nil
	}

	//fmt.Println("Eos")
	//Disp(node)

	canvas := g5.NewCanvas(node.Width, node.Height)

	canvas.Begin()
	{
		if node.CustomRender != nil && !node.CustomRenderTopsKids {
			node.CustomRender(node)
		}

		for kide := node.Kids.Front(); kide != nil; kide = kide.Next() {

			kid := kide.Value.(*Node)

			kidCanvas := renderCanvas(kid)

			kidCanvas.Paint(false, kid.Left , kid.Top, g5.FourOnesFloat32)

			kidCanvas.Free()
		}

		if node.CustomRender != nil && node.CustomRenderTopsKids {
			node.CustomRender(node)
		}
	}
	canvas.End()

	return canvas
}
