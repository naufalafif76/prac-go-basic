package main

import "fmt"

func main() {
	name := "firefly"

	switch name {
	case "kafka":
		fmt.Println("Hello kafka")
	case "firefly":
		fmt.Println("Hello firefly")
	default:
		fmt.Println("Hello")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)

	switch {
	case length > 8:
		fmt.Println("Nama terlalu panjang")
	case length > 2:
		fmt.Println("Nama sudah benar")
	default:
		fmt.Println("Nama terlalu pendek")
	}

}
