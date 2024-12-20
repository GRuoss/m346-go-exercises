package main

import "strconv"

// TODO: implement the function computeGrade
func computerGrade(gotpoints float64, maxPoints float64) string {
	note := (gotpoints/maxPoints)*5 + 1
	if note <= 6 && note >= 1 {
		grade := strconv.FormatFloat(note, 'f', -1, 64)
		return grade
	} else {
		return "error"
	}
}

func main() {
	computerGrade(17.5, 28.0)
}
