package main

import "fmt"

func main() {
	// var names = [3]string{"kafka", "firefly", "wolf"}

	var names [3]string
	names[0] = "silver"
	names[1] = "wolf"
	names[2] = "kafka"
	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [5]int{
		1, 2, 3, 4, 5,
	}
	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))
}
