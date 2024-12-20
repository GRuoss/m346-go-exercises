package main


import (
	"fmt"
)
// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(celsius float64) float64 {
	F:= (celsius * 9/5) + 32
	return F
}
// TODO: implement the function convertFahrenheitToCelsius
func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	C:= (fahrenheit - 32) * 5/9
	return C
}
func main() {
	fmt.Println(convertCelsiusToFahrenheit(10)) // resultate = 50 
	fmt.Println(convertFahrenheitToCelsius(10)) // resultate = -12.222222222222221
}
