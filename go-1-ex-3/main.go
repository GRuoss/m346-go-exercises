package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

func main() {

	var eyes = rand.Intn(5) + 1
	var when = time.Now()
eyesFile, err:= os.Create("eyes.txt") 
if err != nil {
	fmt.Println("Error creating eyes.txt", err)
    return
}
logDice, err := os.Create("dice.log")
if err != nil {
	fmt.Println("Error creating dice.log", err)
    return
}
	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(eyesFile,"the dice shows", eyes, "eyes")

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(logDice,"the dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	// go run ex3/main.go TODO
}
