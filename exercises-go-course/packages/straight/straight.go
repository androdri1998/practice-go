package straight

import "math"

// Ponto refers to a coordinate at cartesian plane
type Point struct {
	X float64
	Y float64
}

func Peccaries(a, b Point) (cx, cy float64) {
	cx = math.Abs(b.X - a.X)
	cy = math.Abs(a.Y - b.Y)
	return
}

func Distance(a, b Point) float64 {
	cx, cy := Peccaries(a, b)

	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
