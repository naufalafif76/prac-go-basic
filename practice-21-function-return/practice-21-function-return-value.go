package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello!"
	}
	return "Hello, " + name + "!"
}

func main() {
	result := getHello("silver")
	fmt.Println(result)
	fmt.Println(getHello(""))
}
