package main

import "fmt"

func main() {
	firstName := "Max"
	lastName := "Mustermann"
	dayOfBirth := 15
	monthOfBirth := 6
	yearOfBirth := 2000
	numberOfSiblings := 2
	heightInMeters := 1.82
	zodiacSign := 'W'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
