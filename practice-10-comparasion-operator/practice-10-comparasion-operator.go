package main

import "fmt"

func main() {
	var name1 = "KAFKA"
	var name2 = "kafka"

	var result bool = name1 == name2 // false
	fmt.Println(result)

	var value1 = 10
	var value2 = 20

	fmt.Println(value1 > value2)  // false
	fmt.Println(value1 < value2)  // true
	fmt.Println(value1 == value2) // false
	fmt.Println(value1 != value2) // true
}
