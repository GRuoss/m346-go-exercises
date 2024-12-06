package main

import "fmt"

type FullName struct {
	firstName string
	lastName  string
}

type Birthdate struct {
	dayOfBirth   int16
	monthOfBirth int16
	yearOfBirth  int16
}

type Profile struct {
	FullName        FullName
	Birthdate       Birthdate
	NumberOfSiblings byte
	ZodiacSign       rune
}

var myName FullName = FullName{
	firstName: "Gregory",
	lastName:  "Ruoss",
}

var myBirthDate Birthdate = Birthdate{
	dayOfBirth:   13,
	monthOfBirth: 11,
	yearOfBirth:  2007,
}

func main() {
	// Profil erstellen
	var me = Profile{
		FullName:        myName,
		Birthdate:       myBirthDate,
		NumberOfSiblings: 3,      
		ZodiacSign: '\u264E', 
	}

	fmt.Println("Profil:", me)

	fmt.Println("Geschwister vor der Geburt:", me.NumberOfSiblings)

	me.NumberOfSiblings++

	fmt.Println("Geschwister nach der Geburt:", me.NumberOfSiblings)
}
