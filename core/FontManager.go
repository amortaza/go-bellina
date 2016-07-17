package core

import (
	"github.com/amortaza/go-g5"
	"github.com/golang/freetype/truetype"
	"strconv"
)

var g_truetypeFontByFontName map[string] *truetype.Font
var g_g4fontByKey map[string] *g5.Gfont

func init() {

	g_truetypeFontByFontName = make(map[string] *truetype.Font)
	g_g4fontByKey = make(map[string] *g5.Gfont)
}

func GetG4Font(fontName string, fontSize int) *g5.Gfont {
	key := fontName + strconv.Itoa(fontSize)

	g4font, ok := g_g4fontByKey[key]

	if !ok {
		truetypeFont, ok2 := g_truetypeFontByFontName[fontName]

		if !ok2 {
			filename := "github.com/amortaza/go-bellina/assets/fonts/" + fontName + ".ttf"

			truetypeFont = g5.LoadTrueTypeFromFile(filename)

			g_truetypeFontByFontName[fontName] = truetypeFont
		}

		g4font = g5.NewGfont(truetypeFont, fontSize)

		g_g4fontByKey[key] = g4font
	}

	return g4font
}
