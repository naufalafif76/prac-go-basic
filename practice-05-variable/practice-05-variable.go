package main

import "fmt"

func main() {
	var name string

	name = "silver wolf"
	fmt.Println(name)

	name = "kafka"
	fmt.Println(name)

	var friendName = "firefly"
	fmt.Println(friendName)

	var age int8 = 26
	fmt.Println(age)

	var friendAge = 22
	fmt.Println(friendAge)

	country := "Hunter"
	fmt.Println(country)

	country = "Express"
	fmt.Println(country)

	var (
		firstName = "Astral"
		lastName  = "Express"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
