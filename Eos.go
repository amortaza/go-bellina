package bl

import (
	"github.com/amortaza/go-g5"
	"github.com/amortaza/go-bellina/core"
	"fmt"
)

func renderCanvas(node *Node) *g5.Canvas {

	if node == nil {
		return nil
	}

	canvas := g5.NewCanvas(node.Width, node.Height)

	canvas.Begin()
	{
		if node.Texture == nil {
			if node.Flags & flag_SKIP_NODE_RECT_DRAW == 0 {
				renderBody(node)
			} else {
				g5.DrawColorRect4v(0, 0, node.Width, node.Height, FourZeroesFloat, FourZeroesFloat, FourZeroesFloat, FourZeroesFloat)
			}
		} else {
			g5.DrawTextureRect(node.Texture, 0,0, node.Width, node.Height, FourOnesFloat)
		}

		if node.CustomRender1_Hook != nil && !node.CustomRender1_TopsLabel {
			node.CustomRender1_Hook(node)
		}

		renderLabel(node)

		if node.CustomRender1_Hook != nil && node.CustomRender1_TopsLabel {
			node.CustomRender1_Hook(node)
		}

		for kide := node.Kids.Front(); kide != nil; kide = kide.Next() {

			kid := kide.Value.(*Node)

			kidCanvas := renderCanvas(kid)

			kidCanvas.Paint(kid.SeeThru, kid.Left, kid.Top, kid.NodeOpacity)

			kidCanvas.Free()

			if kid.BorderTopsCanvas && (kid.Flags & flag_BORDER_ANY != 0) {
				renderBorders(kid.Left, kid.Top, kid)
			}
		}

		if !node.BorderTopsCanvas && (node.Flags & flag_BORDER_ANY != 0) {
			renderBorders(0, 0, node)
		}
	}
	canvas.End()

	return canvas
}

func renderBorders(x, y int, node *Node) {
	//#ifdef nonprod
	if node.BorderThickness == nil {
		fmt.Println("Border flag is set, but border thickness was not defined.")
		return
	}

	var thickness int

	color := []float32{node.BorderRed, node.BorderGreen, node.BorderBlue, 1 }

	if node.Flags & flag_BORDER_LEFT != 0 {
		thickness = node.BorderThickness[0]
		g5.DrawColorRect4v(x, y, thickness, node.Height, color, color, color, color)
	}

	if node.Flags & flag_BORDER_RIGHT != 0 {
		thickness = node.BorderThickness[2]
		g5.DrawColorRect4v(x + node.Width - thickness, y, thickness, node.Height, color, color, color, color)

	}

	if node.Flags & flag_BORDER_TOP != 0 {
		thickness = node.BorderThickness[1]

		g5.DrawColorRect4v(x, y, node.Width, thickness, color, color, color, color)

	}

	if node.Flags & flag_BORDER_BOTTOM != 0 {
		thickness = node.BorderThickness[3]

		g5.DrawColorRect4v(x, y + node.Height - thickness, node.Width, thickness, color, color, color, color)

	}
}

func renderBody(node *Node) {
	color1 := []float32{node.Red1, node.Green1, node.Blue1, 1}

	if node.Flags & flag_COLOR_SOLID != 0 {
		g5.DrawColorRect4v(0, 0, node.Width, node.Height, color1, color1, color1, color1)

	} else if node.Flags & flag_COLOR_HORIflag_GRADIENT != 0 {
		color2 := []float32{node.Red2, node.Green2, node.Blue2, 1}
		g5.DrawColorRect4v(0, 0, node.Width, node.Height, color1, color2, color2, color1)

	} else if node.Flags & flag_COLOR_VERT_GRADIENT != 0 {
		color2 := []float32{node.Red2, node.Green2, node.Blue2, 1}
		g5.DrawColorRect4v(0, 0, node.Width, node.Height, color1, color1, color2, color2)

	} else {
		panic("Unrecognized flag in eos.Render.")
	}
}

func renderLabel(node *Node) {
 	if node.Label == "" {
		return
	}

	g4font := core.GetG4Font(node.FontName, node.FontSize)

	stringTexture := g5.NewStringTexture(node.Label, g4font)
	defer stringTexture.Free()

	fontColor := []float32{node.FontRed, node.FontGreen, node.FontBlue}
	bgColor := []float32{node.Red1, node.Green1, node.Blue1}

	// defaults to 0 which is LABEL_ALIGN_LEFT / LABEL_ALIGN_TOP
	var left, top int

	if node.Flags & flag_LABEL_ALIGN_HCENTER != 0 {
		left = (node.Width - stringTexture.Texture.Width) / 2

	} else if node.Flags & flag_LABEL_ALIGN_RIGHT != 0{
		left = node.Width - stringTexture.Texture.Width
	}

	if node.Flags & flag_LABEL_ALIGN_VCENTER != 0 {
		top = (node.Height - stringTexture.Texture.Height) / 2

	} else if node.Flags & flag_LABEL_ALIGN_BOTTOM != 0 {
		top = node.Height - stringTexture.Texture.Height
	}

	left += node.FontNudgeX
	top += node.FontNudgeY

	g5.DrawStringRect(stringTexture, left, top, fontColor, bgColor, node.LabelOpacity )
}