package avatar

import (
	"fmt"
	"io"

	svg "github.com/ajstarks/svgo"
)

type ShapeType string

const (
	Circle ShapeType = "circle"
	Square ShapeType = "square"
)

type (
	Avatar struct {
		text    string
		options Options
	}
	Options struct {
		BackgroundColor ColorOption
		PrimaryColor    ColorOption
		BorderColor     ColorOption
		FontColor       ColorOption
		Shape           ShapeType
		Size            int
	}
)

func NewAvatar(text string, o Options) Avatar {
	return Avatar{
		options: o,
		text:    text,
	}
}

func (a *Avatar) Generate(w io.Writer) error {
	switch a.options.Shape {
	case Square:
		generateSquare(a.text, a.options, w)
	default:
		generateCircle(a.text, a.options, w)
	}

	return nil
}

func generateSquare(text string, o Options, w io.Writer) {
	size := o.Size
	fontSize := o.Size / 2
	strokeWidth := o.Size / 10

	borderColor, backgroundColor, fontColor := GenerateComplementaryColors(o.BorderColor, o.BackgroundColor, o.FontColor)

	if len(text) > 2 {
		text = text[:2]
	}

	trueSize := size + strokeWidth

	canvas := svg.New(w)
	canvas.Start(trueSize, trueSize)

	squareStyle := fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%d;", backgroundColor, borderColor, strokeWidth)

	canvas.Rect(strokeWidth/2, strokeWidth/2, size, size, squareStyle)

	textX := trueSize / 2
	textY := trueSize/2 + fontSize/3

	textStyle := fmt.Sprintf("text-anchor:middle;font-size:%dpx;font-family:sans-serif;fill:%s;", fontSize, fontColor)
	canvas.Text(textX, textY, text, textStyle)

	canvas.End()
}

func generateCircle(text string, o Options, w io.Writer) {
	width, height := o.Size, o.Size
	circleRadius := o.Size / 2
	fontSize := o.Size / 2
	strokeWidth := o.Size / 10

	borderColor, backgroundColor, fontColor := GenerateComplementaryColors(o.BorderColor, o.BackgroundColor, o.FontColor)

	if len(text) > 2 {
		text = text[:2]
	}

	trueWidth, trueHeight := width+strokeWidth, height+strokeWidth

	canvas := svg.New(w)
	canvas.Start(trueWidth, trueHeight)

	circleStyle := fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%d;", backgroundColor, borderColor, strokeWidth)

	canvas.Circle(trueWidth/2, trueHeight/2, circleRadius, circleStyle)

	textX := trueWidth / 2
	textY := trueHeight/2 + fontSize/3

	textStyle := fmt.Sprintf("text-anchor:middle;font-size:%dpx;font-family:sans-serif;fill:%s;", fontSize, fontColor)
	canvas.Text(textX, textY, text, textStyle)

	canvas.End()
}
