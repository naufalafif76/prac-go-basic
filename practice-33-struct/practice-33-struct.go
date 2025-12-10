package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var silver Customer = Customer{
		Name:    "Silver",
		Address: "Stellaron Hunter",
		Age:     18,
	}

	fmt.Println("Nama:", silver.Name)
	silver.sayHello("Kafka")

	kafka := Customer{
		Name:    "Kafka",
		Address: "Stellaron Hunter",
		Age:     24,
	}

	fmt.Println(kafka)
	kafka.sayHello("Firefly")

	firefly := Customer{"Firefly", "Stellaron Hunter", 20}

	fmt.Println("Nama:", firefly.Name)
	fmt.Println("Umur:", firefly.Age)
}
