/*
Authors: Brijesh Patel, Sudhamsh Mondrati, Ajay Shankar, Aiden James
*/
package main

import (
	"errors"
	"fmt"
	"strconv"
)

var num1 int
var num2 int

func divideByZeroError() error {
	return errors.New("cannot divide by zero")
}
func notInRangeError() error {
	return errors.New("select a choice from 1-6 please")
}

func notIntegerOrFloat() error {
	return errors.New("please only enter an integer")
}

func addValues(a, b int) int {
	return a + b
}

func substractValues(a, b int) int {
	return a - b
}

func muiltplyValues(a, b int) int {
	return a * b
}

func divideValue(a, b int) (float32, error) {
	if float32(b) == 0 {
		return -1, divideByZeroError()
	} else {
		return float32(a) / float32(b), nil
	}
}

func reScanInput(a int) (int, error) {
	if a > 0 && a < 7 {
		return a, nil
	} else {
		return -1, notInRangeError()
	}
}

func moduloValue(a, b int) int {
	return a % b
}

func ParseIntFromInput() int {
  var input string
  for true {
    fmt.Scan(&input)
		if n1, err := strconv.ParseInt(input, 10, 64); err == nil {
			return int(n1)
		} else {
			fmt.Println("please input an integer value")
		}
	}
  return 0
}

func getInputValues() (int, int) {
	num1 = 0
	num2 = 0

  
  fmt.Println("Please enter the first value:")
  
	num1 = ParseIntFromInput()
  fmt.Println("Please enter the second value:")
	num2 = ParseIntFromInput()
	
	return num1, num2
}

func main() {
	//
	var condTrue int
	condTrue = 1
	for condTrue == 1 {

		fmt.Println("Enter a number what calculation you want to have done:")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Modulo")
		fmt.Println("6. QUIT")

		var calculationChoice int
		calculationChoice = ParseIntFromInput()


		fmt.Println()

		switch calculationChoice {
		case 1:
			num1, num2 = getInputValues()
			fmt.Println("The values are added: ", addValues(num1, num2))
		case 2:
			num1, num2 = getInputValues()
			fmt.Println("The values are subtracted: ", substractValues(num1, num2))
		case 3:
			num1, num2 = getInputValues()
			fmt.Println("The values are muitlplied: ", muiltplyValues(num1, num2))
		case 4:
			num1, num2 = getInputValues()
			resultFromMethod, errorValue := divideValue(num1, num2)
			if errorValue == nil {
				fmt.Println("The values are divided: ", resultFromMethod)
			} else {
				fmt.Println(errorValue)
			}
		case 5:
			num1, num2 = getInputValues()
			fmt.Println("The values are modulo: ", moduloValue(num1, num2))
		case 6:
			condTrue = condTrue + 1
    default:
      fmt.Println("Please enter an integer from 1 to 6")
		}
	}

	fmt.Println("Quitting now, thank you.")

} // End of the main method
