package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 15, 25, 30)
	fmt.Println(total)

	// memasukkan slice ke dalam variadic function
	slices := []int{16, 32, 22}
	numbers := sumAll(slices...)
	fmt.Println(numbers)
}
