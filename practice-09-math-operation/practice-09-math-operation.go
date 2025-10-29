package main

import "fmt"

func main() {
	// binary operator
	// a + b
	// a - b
	// a * b
	// a / b
	// a % b
	// a & b
	// a | b
	// a ^ b
	// a << b
	// a >> b

	var result = 10 + 5
	fmt.Println(result)

	var a = 10
	var b = 5
	var c = a * b
	fmt.Println(c)

	// augmented operator
	// a += 10
	// a -= 10
	// a *= 10
	// a /= 10
	// a %= 10
	// a &= 10
	// a |= 10
	// a ^= 10
	// a <<= 10
	// a >>= 10

	var i = 10
	i += 10
	fmt.Println(i)

	// unary operator
	// a++
	// a--
	// ++a
	// --a

	i++ // i = i + 1
	fmt.Println(i)

	var negative = -10
	var positive = +10
	fmt.Println(negative)
	fmt.Println(positive)
}
