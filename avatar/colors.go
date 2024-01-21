package avatar

import (
	"fmt"
	"math/rand"
	"time"
)

type ColorOption struct {
	Color string
	Set   bool
}

func NewColorOption(color string) ColorOption {
	return ColorOption{Color: color, Set: true}
}

func hexToRGB(hex string) (int, int, int) {
	var r, g, b int
	_, err := fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
	if err != nil {
		r, g, b = 0, 0, 0
	}
	return r, g, b
}

func adjustColor(r, g, b int, factor float64) (int, int, int) {
	newR := int(float64(r) * factor)
	newG := int(float64(g) * factor)
	newB := int(float64(b) * factor)

	newR = min(max(newR, 0), 255)
	newG = min(max(newG, 0), 255)
	newB = min(max(newB, 0), 255)

	return newR, newG, newB
}

// rgbToHex converts RGB values to a hex color string.
func rgbToHex(r, g, b int) string {
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

// complement calculates the complementary color.
func complement(r, g, b int) (int, int, int) {
	return 255 - r, 255 - g, 255 - b
}

func GenerateColors(border, background, font ColorOption) (borderColor, backgroundColor, fontColor string) {
	goodColors := []string{"#0066cc", "#ffcc00", "#ff0099", "#33cc33", "#9933ff", "#ff6666"}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	defaultColor := goodColors[rng.Intn(len(goodColors))]

	var baseColor string
	if border.Set {
		baseColor = border.Color
	} else if background.Set {
		baseColor = background.Color
	} else if font.Set {
		baseColor = font.Color
	} else {
		baseColor = defaultColor
	}

	r, g, b := hexToRGB(baseColor)

	backgroundColor = background.Color
	br, bg, bb := hexToRGB(backgroundColor)
	if !background.Set {
		br, bg, bb = int(min(255, float64(r)*1.2)), int(min(255, float64(g)*1.2)), int(min(255, float64(b)*1.2))
		backgroundColor = rgbToHex(br, bg, bb)
	}

	borderColor = border.Color
	if !border.Set {
		borderColorR, borderColorG, borderColorB := adjustColor(br, bg, bb, 0.75)
		borderColor = rgbToHex(borderColorR, borderColorG, borderColorB)
	}

	fontColor = font.Color
	if !font.Set {
		fr, fg, fb := adjustColor(br, bg, bb, 0.45)
		fontColor = rgbToHex(fr, fg, fb)
	}

	return
}
