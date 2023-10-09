package calculate

import "math"

func GetRectangleArea(width, height float32) float32 {
	return width * height
}

func GetCircleArea(radius float32) float32 {
	return math.Pi * radius * radius
}
