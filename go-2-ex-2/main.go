package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

	fibs[2] = fibs[0] + fibs[1]
	fibs[3] = fibs[1] + fibs[2]
	fibs[4] = fibs[2] + fibs[3]

	nextFib := fibs[3] + fibs[4]
	fibs = append(fibs, nextFib) 

	for i := 0; i < 3; i++ {
		nextFib = fibs[len(fibs)-2] + fibs[len(fibs)-1]
		fibs = append(fibs, nextFib)
	}
	fmt.Println(fibs) // expected output: [1 1 2 3 5 8 13 21 34]
}
