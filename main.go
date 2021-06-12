package main

import "digitalHouse/samples"

func main() {
	john:= samples.NewPerson("John", "Doe",31)
	john.Eat()
	john.Sleep()
	john.TurnYears()
}
