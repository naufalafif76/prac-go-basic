package main

import (
	"fmt"
)

// alias untuk function filter
type filterFunc func(string) string

func sayHelloWithFilter(name string, filter filterFunc) string {
	nameFiltered := filter(name)
	return "Hello " + nameFiltered
}

func filterName(name string) string {
	if len(name) > 9 {
		return "nama terlalu panjang"
	} else {
		return name
	}
}

func main() {
	nameFiltered := filterName
	helloName := sayHelloWithFilter("silver", nameFiltered)

	fmt.Println(helloName)
}
