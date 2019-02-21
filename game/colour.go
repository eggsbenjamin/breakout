package game

type Colour struct {
	R, G, B, A uint8
}

func NewColour(r, g, b, a uint8) Colour {
	return Colour{
		r, g, b, a,
	}
}

var (
	ColourBlack = NewColour(0, 0, 0, 0)
)
