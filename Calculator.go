package main

import (
	"fmt"
)

func main() {
	var option string

	fmt.Println("\x1b[34m=====================================================\x1b[0m")
	fmt.Println("           \x1b[35mCalculator Application\x1b[0m")
	fmt.Println("              \x1b[35mFinal Project\x1b[0m                   ")
	fmt.Println("  \x1b[35mCreated By: Raffaele Cristallo (2031354)\x1b[0m")
	fmt.Println("\x1b[34m=====================================================\x1b[0m")

	fmt.Println("\n\x1b[35m| If you'd like to quit the calculator application type '/quit'\x1b[0m")

	for {
		var num1, num2 float64

		fmt.Println("\nEnter an option listed: (+, -, *, /, or '/quit' to quit:)")
		fmt.Scanln(&option)

		if option == "/quit" {
			fmt.Println("\x1b[31m=====================================================\x1b[0m")
			fmt.Println("              \x1b[35mFinal Project\x1b[0m                   ")
			fmt.Println("  \x1b[35mCreated By: Raffaele Cristallo (2031354)\x1b[0m")
			fmt.Println("\x1b[31m=====================================================\x1b[0m")
			return
		}

		//first user input number
		fmt.Println("Enter your first number: ")
		num1, err := getValidInput()
		if err != nil {
			printError(err.Error())
			continue
		}

		//second user input number
		fmt.Println("Enter your second number: ")
		num2, err = getValidInput()
		if err != nil {
			printError(err.Error())
			continue
		}

		result := 0.0
		switch option {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 != 0 { //this checks if the user is trying to divide by 0
				result = num1 / num2
			} else {
				printError("\x1b[31mCannot divide by zero\x1b[0m") //User cannot divide by 0 or else will get an error
				continue
			}
		default:
			printError("\x1b[31mYou've chosen an invalid option\x1b[0m") //Error messaged displayed if invalid option
			continue
		}
		fmt.Println("Result:", num1, option, num2, "=", result) //Prints the result
	}
}

func getValidInput() (float64, error) {
	var input string
	fmt.Scanln(&input)
	num, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	return num, nil
}

// this function uses a string as input and returns a float
// if the input is valid, or return an error
func parseInput(input string) (float64, error) {
	var num float64
	_, err := fmt.Sscanf(input, "%f", &num)
	if err != nil {
		return 0, fmt.Errorf("\x1b[31m%v is not a valid number\x1b[0m", input)
	}
	return num, nil
}

func printError(message string) {
	fmt.Println("\x1b[31mError:\x1b[0m", message) //Prints an error in the console in a red colour.
}
