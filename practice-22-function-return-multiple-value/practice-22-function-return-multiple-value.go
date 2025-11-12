package main

import "fmt"

func getFullName(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

func main() {
	firstName, _ := getFullName("silver", "wolf")
	fmt.Println(firstName)
}
