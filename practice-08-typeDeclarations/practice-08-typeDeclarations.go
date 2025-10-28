package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpSilver NoKTP = "3325xxx3xxx2xxx1"
	var marriedStatus Married = false

	fmt.Println(noKtpSilver)
	fmt.Println(NoKTP("6645xxx3xxx2xxx1"))
	fmt.Println(marriedStatus)
}
