package main

import "math"

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
func computeHypotenuse(a float64, b float64) float64 {
	c := math.Sqrt( math.Pow(a, 2) + math.Pow(b, 2))
	return c
}
func main() {
	// TODO: call the function computeHypotenuse
	computeHypotenuse(3.0, 4.0)
}
