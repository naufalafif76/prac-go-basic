package main

import "fmt"

func getFullMember() (firstMember string, secondMember string, lastMember string) {
	firstMember = "kafka"
	secondMember = "firefly"
	lastMember = "silver"
	return // tidak perlu deklarasikan variabel lagi, karena sudah dideklarasikan di atas
}

func main() {
	first, second, last := getFullMember()
	fmt.Println(first+"\n", second+"\n", last)
}
