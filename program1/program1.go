/*
	Author: Sudhamsh Mondrati, Melvin Priyemskiy
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to first program.")
	// creating numbers, strings & booleans
	defaultNumber16 := 16
	var number18 int = 18

	var stringOne string = "String one."
	stringTwo := "String two."

	var floatNumber float32 = 3942.49502
	floatPi := 3.14159265358979

	var defaultBoolean bool = false
	trueBoolean := true

	// printing above info
	fmt.Printf("Numbers: %v, %v\n", defaultNumber16, number18)
	fmt.Printf("Strings: %v, %v\n", stringOne, stringTwo)
	fmt.Printf("Floats: %f, %f\n", floatNumber, floatPi)
	fmt.Printf("Booleans: %t, %t\n", defaultBoolean, trueBoolean)

	//taking the sum
	fmt.Printf("\nSum of Numbers: %v\n", sum(defaultNumber16, number18))
	fmt.Printf("Sum of Floats: %v\n", sum_float(floatNumber, floatPi))

	//taking the difference
	fmt.Printf("\nDifference of Numbers: %v\n", difference(defaultNumber16, number18))
	fmt.Printf("Difference of Floats: %v\n", diff_float(floatNumber, floatPi))

	// replacing strings
	otherString := string_replace(stringOne, stringTwo)
	fmt.Printf("\nAfter replacing Strings: %v\n", otherString)

	//adding strings
	var combinedString string = string_conc(stringOne, stringTwo)
	fmt.Printf("After combining Strings: %v\n", combinedString)

	//converting using cast
	fmt.Printf("\nConverting pi to int: %d\n", int(floatPi))
	fmt.Printf("Converting 18 to float32: %0.2f\n", float32(number18))

	//using boolean functions
	fmt.Printf("\nInverting given boolean value (false to true): %t\n", inverseBool(defaultBoolean))
	fmt.Printf("Applying bitwise Or to the given boolean values (false || true): %t\n", bitwiseOr(defaultBoolean, trueBoolean))
}

func sum(num1 int, num2 int) int {
	var result int = num1 + num2
	return result
}

func difference(num1 int, num2 int) int {
	var result int = num1 - num2
	return result
}

func sum_float(num1 float32, num2 float64) float64 {
	return float64(num1) + num2
}

func diff_float(num1 float32, num2 float64) float64 {
	return float64(num1) - num2
}

func string_replace(str1 string, str2 string) string {
	newString := strings.Replace(str1, str1, str2, -1)
	return newString
}

func string_conc(str1 string, str2 string) string {
	return str1 + str2
}

func inverseBool(userInput bool) bool {
	return !userInput
}

func bitwiseOr(input1 bool, input2 bool) bool {
	if input1 || input2 {
		return true
	}
	return false
}
