package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	for count := 1; count <= 5; count++ {
		fmt.Println("hitungan ke", count)
	}

	slice := []string{"silver", "wolf", "kafka", "firefly"}
	for i, v := range slice {
		// fmt.Println(slice[i])
		fmt.Println("Index ke", i, "=", v) // gunakan _ untuk pengganti index yang tidak digunakan
	}

	person := map[string]string{"name": "Silver", "job": "Programmer", "affiliation": "Hunter"}
	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
