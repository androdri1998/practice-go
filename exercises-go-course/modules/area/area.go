package area

import "math"

const PI = 3.14

func Circle(ray float64) float64 {
	return math.Pow(ray, 2) * PI
}

func Rect(base, height float64) float64 {
	return base * height
}
