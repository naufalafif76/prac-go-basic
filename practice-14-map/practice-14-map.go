package main

import "fmt"

func main() {
	person := map[string]string{
		"name":        "Silver",
		"affiliation": "Hunter",
	}

	person["job"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["job"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Silver"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}
