package main

import "fmt"

func main() {
	const firstName string = "silver"
	const lastName = "wolf"
	const value = 4

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	const (
		name1  string = "kafka"
		name2         = "firefly"
		value1 int    = 2
	)

}
