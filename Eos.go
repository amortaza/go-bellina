package bl

import (
	"g4"
	"bellina/core"
	"fmt"
)

func renderCanvas(node *Node) *g4.Canvas {

	canvas := g4.NewCanvas(node.Width, node.Height)

	canvas.Begin()
	{
		if node.Texture == nil {
			renderBody(node)
		} else {
			g4.DrawTextureRect(node.Texture, 0,0, node.Width, node.Height, FourOnesFloat)
		}

		renderLabel(node)

		for kide := node.Kids.Front(); kide != nil; kide = kide.Next() {

			kid := kide.Value.(*Node)

			kidCanvas := renderCanvas(kid)

			kidCanvas.Paint(kid.SeeThru, kid.Left, kid.Top, kid.NodeOpacity)

			kidCanvas.Free()

			if kid.BorderTopsCanvas && (kid.Flags & Z_BORDER_ANY != 0) {
				renderBorders(kid.Left, kid.Top, kid)
			}
		}

		if !node.BorderTopsCanvas && (node.Flags & Z_BORDER_ANY != 0) {
			renderBorders(0, 0, node)
		}
	}
	canvas.End()

	return canvas
}

func renderBorders(x, y int32, node *Node) {
	//#ifdef nonprod
	if node.BorderThickness == nil {
		fmt.Println("Border flag is set, but border thickness was not defined.")
		return
	}

	var thickness int32

	color := []float32{node.BorderRed, node.BorderGreen, node.BorderBlue, 1 }

	if node.Flags & Z_BORDER_LEFT != 0 {
		thickness = node.BorderThickness[0]
		g4.DrawColorRect(x, y, thickness, node.Height, color, color, color, color)
	}

	if node.Flags & Z_BORDER_RIGHT != 0 {
		thickness = node.BorderThickness[2]
		g4.DrawColorRect(x + node.Width - thickness, y, thickness, node.Height, color, color, color, color)

	}

	if node.Flags & Z_BORDER_TOP != 0 {
		thickness = node.BorderThickness[1]

		g4.DrawColorRect(x, y, node.Width, thickness, color, color, color, color)

	}

	if node.Flags & Z_BORDER_BOTTOM != 0 {
		thickness = node.BorderThickness[3]

		g4.DrawColorRect(x, y + node.Height - thickness, node.Width, thickness, color, color, color, color)

	}
}

func renderBody(node *Node) {
	color1 := []float32{node.Red1, node.Green1, node.Blue1, 1}

	if node.Flags & Z_COLOR_SOLID != 0 {
		g4.DrawColorRect(0, 0, node.Width, node.Height, color1, color1, color1, color1)

	} else if node.Flags & Z_COLOR_HORIZ_GRADIENT != 0 {
		color2 := []float32{node.Red2, node.Green2, node.Blue2, 1}
		g4.DrawColorRect(0, 0, node.Width, node.Height, color1, color2, color2, color1)

	} else if node.Flags & Z_COLOR_VERT_GRADIENT != 0 {
		color2 := []float32{node.Red2, node.Green2, node.Blue2, 1}
		g4.DrawColorRect(0, 0, node.Width, node.Height, color1, color1, color2, color2)

	} else {
		panic("Unrecognized flag in eos.Render.")
	}
}

func renderLabel(node *Node) {
 	if node.Label == "" {
		return
	}

	g4font := core.GetG4Font(node.FontName, node.FontSize)

	stringTexture := g4.NewStringTexture(node.Label, g4font)
	defer stringTexture.Free()

	fontColor := []float32{node.FontRed, node.FontGreen, node.FontBlue}
	bgColor := []float32{node.Red1, node.Green1, node.Blue1}

	// defaults to 0 which is LABEL_ALIGN_LEFT / LABEL_ALIGN_TOP
	var left, top int32

	if node.Flags & Z_LABEL_ALIGN_HCENTER != 0 {
		left = (node.Width - stringTexture.Texture.Width) / 2

	} else if node.Flags & Z_LABEL_ALIGN_RIGHT != 0{
		left = node.Width - stringTexture.Texture.Width
	}

	if node.Flags & Z_LABEL_ALIGN_VCENTER != 0 {
		top = (node.Height - stringTexture.Texture.Height) / 2

	} else if node.Flags & Z_LABEL_ALIGN_BOTTOM != 0 {
		top = node.Height - stringTexture.Texture.Height
	}

	left += node.FontNudgeX
	top += node.FontNudgeY

	g4.DrawStringRect(stringTexture, left, top, fontColor, bgColor, node.LabelOpacity )
}