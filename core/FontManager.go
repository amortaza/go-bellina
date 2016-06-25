package core

import (
	"github.com/amortaza/go-g4"
	"github.com/golang/freetype/truetype"
	"strconv"
)

var g_truetypeFontByFontName map[string] *truetype.Font
var g_g4fontByKey map[string] *g4.G4Font

func init() {

	g_truetypeFontByFontName = make(map[string] *truetype.Font)
	g_g4fontByKey = make(map[string] *g4.G4Font)
}

func GetG4Font(fontName string, fontSize int32) *g4.G4Font {
	key := fontName + strconv.Itoa(int(fontSize))

	g4font, ok := g_g4fontByKey[key]

	if !ok {
		truetypeFont, ok2 := g_truetypeFontByFontName[fontName]

		if !ok2 {
			filename := "assets/fonts/" + fontName + ".ttf"

			truetypeFont = g4.LoadTrueTypeFromFile(filename)

			g_truetypeFontByFontName[fontName] = truetypeFont
		}

		g4font = g4.NewG4Font(truetypeFont, fontSize)

		g_g4fontByKey[key] = g4font
	}

	return g4font
}
