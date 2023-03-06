package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to first program.")
	// creating numbers, strings & booleans
	defaultNumber16 := 16
	var number18 int = 127

	var stringOne string = "String one."
	stringTwo := "String two."

	var floatNumber float32 = 391042.49502
	floatPi := 3.14159265358979

	var defaultBoolean bool = false
	trueBoolean := true

	// printing above info
	fmt.Printf("Numbers: %v, %v\n", defaultNumber16, number18)
	fmt.Printf("Strings: %v, %v\n", stringOne, stringTwo)
	fmt.Printf("Floats: %f, %f\n", floatNumber, floatPi)
	fmt.Printf("Booleans: %t, %t\n", defaultBoolean, trueBoolean)

	fmt.Printf("Sum of Numbers: %v\n", sum(defaultNumber16, number18))

}

func sum(num1 int, num2 int) int {
	var result int = num1 + num2
	return result
}
