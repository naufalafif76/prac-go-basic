package main

import "fmt"

func main() {
	namaDepan := "silver"
	sayHelloTo(namaDepan, "wolf")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
