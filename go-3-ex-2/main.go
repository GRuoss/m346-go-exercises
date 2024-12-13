package main

import "fmt"

const (
	Aries       = '\u2648' // Widder
	Taurus      = '\u2649' // Stier
	Gemini      = '\u264a' // Zwillinge
	Cancer      = '\u264b' // Krebs
	Leo         = '\u264c' // Löwe
	Virgo       = '\u264d' // Jungfrau
	Libra       = '\u264e' // Waage
	Scorpius    = '\u264f' // Skorpion
	Sagittarius = '\u2650' // Schütze
	Capricornus = '\u2651' // Steinbock
	Aquarius    = '\u2652' // Wassermann
	Pisces      = '\u2653' // Fische
)

func outputDateRange(zodiacSign rune) {
	fmt.Printf("%c: ", zodiacSign)

	switch{
		case zodiacSign == Aries: fmt.Println("21.03. - 19.04") // March 21 - April 19
		case zodiacSign == Taurus: fmt.Println("20.04. - 20.05") // March 21 - April 19
		case zodiacSign == Gemini: fmt.Println("21.05. - 21.06") // March 21 - April 19
		case zodiacSign == Cancer: fmt.Println("22.06. - 22.07") // March 21 - April 19
		case zodiacSign == Leo: fmt.Println("23.07. - 22.08") // March 21 - April 19
		case zodiacSign == Virgo: fmt.Println("23.08. - 22.09") // March 21 - April 19
		case zodiacSign == Libra: fmt.Println("23.09. - 23.10") // March 21 - April 19
		case zodiacSign == Scorpius: fmt.Println("24.10. - 21.11") // March 21 - April 19
		case zodiacSign == Sagittarius: fmt.Println("22.11. - 21.12") // March 21 - April 19
		case zodiacSign == Capricornus: fmt.Println("22.12. - 19.01")// March 21 - April 19
		case zodiacSign == Aquarius: fmt.Println("20.1. - 18.02") // March 21 - April 19
		case zodiacSign == Pisces: fmt.Println("19.2. - 20.03") // March 21 - April 19
		default: fmt.Println("unknown")
		
	}
	// TODO: Replace if, else if branching with switch/case.
	// TODO: Define all 12 cases...
	// TODO: ...and consider a default case.
}

func main() {
	for zodiacSign := Aries; zodiacSign <= Pisces; zodiacSign++ {
		outputDateRange(zodiacSign)
	}
}
