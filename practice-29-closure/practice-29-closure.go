package main

import "fmt"

func main() {
	name := "silver"
	counter := 0
	increment := func() {
		name := "wolf"
		fmt.Println(name)
		counter++
	}

	increment()

	fmt.Println(name)
	fmt.Println(counter)
}
