package main

import "fmt"
import "errors"

var num1 int
var num2 int

func dividebyZeroError() error {
	return errors.New("Cannot divide by zero")
}
func notInRangeError() error {
	return errors.New("Select a choice from 1-6 please")
} 

func notIntegerOrFloat() error {
	return errors.New("The value should be a integer and if you entered zero do the math yourself")
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


func getInputValues() (int, int) {
	num1 = 0
	num2 = 0
	fmt.Println("Please enter the first value:")
	fmt.Scan(&num1)

	var calculationChoiceInRange int
	calculationChoiceInRange = 1
	for calculationChoiceInRange == 1 {
		_, errorValue := isInt(num1) 
		if errorValue != nil {
			fmt.Println(errorValue)
			fmt.Println("Please enter the first value:")
			fmt.Scan(&num1)
		} else {
			calculationChoiceInRange = calculationChoiceInRange + 2
		}
	}

	fmt.Println("Please enter the second value:")
	fmt.Scan(&num2)
	var calculationChoiceInRange2 int
	calculationChoiceInRange2 = 1
	for calculationChoiceInRange2 == 1 {
		_, errorValue := isInt(num2) 
		if errorValue != nil {
			fmt.Println(errorValue)
			fmt.Println("Please enter the second value:")
			fmt.Scan(&num2)
		} else {
			calculationChoiceInRange2 = calculationChoiceInRange2 + 2
		}
	}

	return num1, num2
}

func isInt(a int) (int, error) {
	if a == 0 {
		return -1, notIntegerOrFloat()
	} else {
		return a, nil
	}
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
		fmt.Scan(&calculationChoice)

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
		}
	}
	
	fmt.Println("Quitting now, thank you.")
	
} // End of the main method
