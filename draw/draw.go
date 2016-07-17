package bl_draw

import "github.com/amortaza/go-g4"

func DrawPill(left, top, width, height int, rgba []float32) {
	radius := height / 2

	g4.DrawRingOpen(left + radius, top + radius, radius, radius, 90, 270, rgba[3], rgba[3], rgba)
	g4.DrawColorRect4v(left + radius,top, width-radius-radius, height, rgba,rgba,rgba,rgba)
	g4.DrawRingOpen(left + width - radius, top + radius, radius, radius, 270, 90, rgba[3], rgba[3], rgba)
}

func DrawRoundedRect(left, top, width, height, radius int, rgba []float32) {

	x0 := left + radius
	w, h := width - 2 * radius, height - 2 * radius
	x1 := x0 + w
	y0 := top + radius
	y1 := y0 + h

	g4.DrawRingOpen(x0, y0, radius, radius, 180, 270, rgba[3], rgba[3], rgba)
	g4.DrawRingOpen(x1, y0, radius, radius, 270, 180, rgba[3], rgba[3], rgba)
	g4.DrawRingOpen(x1, y1, radius, radius, 0, 90, rgba[3], rgba[3], rgba)
	g4.DrawRingOpen(x0, y1, radius, radius, 90, 180, rgba[3], rgba[3], rgba)
	g4.DrawColorRect4v(x0, top, w, height, rgba,rgba,rgba,rgba)
	g4.DrawColorRect4v(left, y0, radius, h, rgba,rgba,rgba,rgba)
	g4.DrawColorRect4v(x1, y0, radius, h, rgba,rgba,rgba,rgba)
}
