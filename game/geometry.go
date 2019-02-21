package game

type Rectangle struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func Intersects(r1, r2 Rectangle) bool {
	if r1.X > r2.X+r2.Width || r2.X > r1.X+r1.Width {
		return false
	}

	if r1.Y+r1.Height < r2.Y || r2.Y+r2.Height < r1.Y {
		return false
	}

	return true
}

type Point struct {
	X int32
	Y int32
}

func NewPoint(x, y int32) Point {
	return Point{
		X: x,
		Y: y,
	}
}

type Line struct {
	Slope float64
	Yint  float64
}

func NewLine(a, b Point) Line {
	Slope := float64(b.Y-a.Y) / float64(b.X-a.X)
	yint := float64(a.Y) - Slope*float64(a.X)
	return Line{Slope, yint}
}

func EvalX(l Line, x float64) float64 {
	return l.Slope*x + l.Yint
}

func Intersection(l1, l2 Line) (Point, bool) {
	if l1.Slope == l2.Slope {
		return Point{}, false
	}
	x := (l2.Yint - l1.Yint) / (l1.Slope - l2.Slope)
	y := EvalX(l1, x)
	return Point{int32(x), int32(y)}, true
}
