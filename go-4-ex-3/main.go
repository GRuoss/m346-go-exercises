package main

import (
	"errors"
	"fmt"
	"math"
)

func computeQuadraticFormula(a float64, b float64, c float64) ([]float64, error) {
	D := math.Pow(b, 2) - 4*a*c
	if D < 0 {
		return nil, errors.New("error: D < 0")
	} else if D == 0 {
		x := (-b + math.Sqrt(D)) / (2 * a)
		return []float64{x}, nil
	} else {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		return []float64{x1, x2}, nil
	}
}

func main() {
	
	solutions, err := computeQuadraticFormula(1.0, -3.0, 2.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Lösungen für a=1, b=-3, c=2:", solutions)
	}

	solutions, err = computeQuadraticFormula(1.0, 2.0, 1.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Lösungen für a=1, b=2, c=1:", solutions) 
	}

	solutions, err = computeQuadraticFormula(1.0, 1.0, 1.0)
	if err != nil {
		fmt.Println(err) 
	} else {
		fmt.Println("Lösungen für a=1, b=1, c=1:", solutions)
	}
}
