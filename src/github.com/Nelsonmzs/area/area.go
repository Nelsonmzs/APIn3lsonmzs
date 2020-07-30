package area

import "math"

const Pi = 3.1416

// Circ
func Circ(raio float64) float64  {
	return math.Pow(raio, 2) * Pi
}
