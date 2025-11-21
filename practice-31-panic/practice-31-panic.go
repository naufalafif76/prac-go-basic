package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
	message := recover()
	if message != nil {
		fmt.Println("Error message:", message)
	}
}

func runApp(error bool) {
	defer endApp()
	fmt.Println("Aplikasi berjalan")
	if error {
		panic("Aplikasi Error!")
	}
}

func main() {
	runApp(false)
}
