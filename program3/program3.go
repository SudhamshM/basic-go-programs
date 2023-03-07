/*
Authors: Brijesh Patel, Sudhamsh Mondrati, Ajay Shankar
*/
package main

import (
	"errors"
	"fmt"
	"strconv"
)

func dividebyZeroError() error {
	return errors.New("cannot divide by zero")
}
func notInRangeError() error {
	return errors.New("select a choice from 1-6 please")
}

func notIntegerOrFloat() error {
	return errors.New("the value should be a integer or float.")
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

func divideValue(a, b int) (int, error) {
	if b == 0 {
		return -1, dividebyZeroError()
	} else {
		return a / b, nil
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

func getInputValues() (int, int, error) {

	fmt.Println("Please enter the first value:")
	var userInput1, userInput2 string
	fmt.Scan(&userInput1)
	if _, err := strconv.ParseInt(userInput1, 10, 64); err == nil {
		// it is an integer
	} else {
		// Try to parse input as a float
		if _, err := strconv.ParseFloat(userInput1, 64); err == nil {
			// it is a float
		} else {
			return -1, -1, notIntegerOrFloat()
		}
	}

	fmt.Println("Please enter the second value:")
	fmt.Scan(&userInput2)

	if _, err := strconv.ParseInt(userInput2, 10, 64); err == nil {
		// it is an integer
	} else {
		// Try to parse input as a float
		if _, err := strconv.ParseFloat(userInput2, 64); err == nil {
			// it is a float
		} else {
			return -1, -1, notIntegerOrFloat()
		}
	}

	return num1, num2, nil
}

/*
func divide(a, b int) int {
	return a / b
}
*/

func main() {
	//
	var condTrue int
	condTrue = 1

	num1, num2 := 0, 0

	for condTrue == 1 {

		// Asking the user to input the size of the array (group)
		fmt.Println("Enter a number what calculation you want to have done:")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Modulo")
		fmt.Println("6. QUIT")
		// Saving the user input to the memory address of groupSize

		var calculationChoice int
		fmt.Scan(&calculationChoice)
		//_, errorValue := reScanInput(calculationChoice)
		//fmt.Println("Error: ", errorValue)
		var calculationChoiceInRange int
		calculationChoiceInRange = 1
		for calculationChoiceInRange == 1 {
			_, errorValue := reScanInput(calculationChoice)
			if errorValue != nil {
				fmt.Println(errorValue)
				fmt.Scan(&calculationChoice)

			} else {
				calculationChoiceInRange = calculationChoiceInRange + 2
				break
			}
		}

		//fmt.Println(num1)
		//fmt.Println(num2)
		//fmt.Println(calculationChoice)
		fmt.Println()
		fmt.Println()

		switch calculationChoice {

		case 1:
			var errorVal error
			num1, num2, errorVal = getInputValues()
			if errorVal == nil {
				fmt.Println("The values are added: ", addValues(num1, num2))
			} else {
				fmt.Println(errorVal)
			}
		case 2:
			var errorVal error
			num1, num2, errorVal = getInputValues()
			if errorVal == nil {
				fmt.Println("The values are subtracted: ", substractValues(num1, num2))
			} else {
				fmt.Println(errorVal)
			}
		case 3:
			var errorVal error
			num1, num2, errorVal = getInputValues()

			if errorVal == nil {
				fmt.Println("The values are muitlplied: ", muiltplyValues(num1, num2))
			} else {
				fmt.Println(errorVal)
			}
		case 4:
			var errorVal error
			num1, num2, errorVal = getInputValues()
			if errorVal == nil {
				resultFromMethod, errorValue := divideValue(num1, num2)
				if errorValue == nil {
					fmt.Println("The values are divided: ", resultFromMethod)
				} else {
					fmt.Println(errorValue)
				}
			} else {
				fmt.Println(errorVal)
			}

		case 5:
			var errorVal error
			num1, num2, errorVal = getInputValues()

			if errorVal == nil {
				fmt.Println("The values are modulo: ", moduloValue(num1, num2))
			} else {
				fmt.Println(errorVal)
			}
		case 6:
			condTrue = condTrue + 1
		}
	}

	fmt.Println("Quitting now, thank you.")

} // End of the main method
