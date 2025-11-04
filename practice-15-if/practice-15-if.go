package main

import "fmt"

func main() {
	name := "firefly"

	if name == "kafka" {
		fmt.Println("Hello Kafka")
	} else if name == "firefly" {
		fmt.Println("Hello firefly")
	} else {
		fmt.Println("Hello World")
	}

	// if length := len(name); length > 5 {
	// 	fmt.Println("Nama terlalu panjang")
	// } else {
	// 	fmt.Println("Nama sudah benar")
	// }
}
