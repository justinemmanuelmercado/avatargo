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

// hexToRGB converts a hex color string to RGB.
func hexToRGB(hex string) (int, int, int) {
	var r, g, b int
	_, err := fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
	if err != nil {
		// Handle errors or set default values
		r, g, b = 0, 0, 0
	}
	return r, g, b
}

// adjustColor creates a variation of the given color.
func adjustColor(r, g, b int, factor float64) (int, int, int) {
	newR := int(float64(r) * factor)
	newG := int(float64(g) * factor)
	newB := int(float64(b) * factor)

	// Ensure RGB values are within valid range
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

// GenerateComplementaryColors generates colors based on the provided or default colors.
func GenerateComplementaryColors(border, background, font ColorOption) (borderColor, backgroundColor, fontColor string) {
	// List of good-looking colors
	goodColors := []string{"#0066cc", "#ffcc00", "#ff0099", "#33cc33", "#9933ff", "#ff6666"}

	// Seed the random number generator
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Select a random color from the list as the default
	defaultColor := goodColors[rng.Intn(len(goodColors))]

	// Determine base color for calculations
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

	// Set colors based on inputs
	borderColor = border.Color
	if !border.Set {
		borderColorR, borderColorG, borderColorB := adjustColor(br, bg, bb, 0.75) // Darken by 15%
		borderColor = rgbToHex(borderColorR, borderColorG, borderColorB)
	}

	fontColor = font.Color
	if !font.Set {
		fr, fg, fb := adjustColor(br, bg, bb, 0.45)
		fontColor = rgbToHex(fr, fg, fb)
	}

	return
}
