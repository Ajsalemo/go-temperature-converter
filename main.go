package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var originUnit string
var originValue float64

var shouldConvertAgain string

var err error

var errInvalidArguments = errors.New("invalid arguments")
var errReadingInput = errors.New("error reading input")

func main() {
	if len(os.Args) != 2 {
		printError(errInvalidArguments)
	}

	originUnit = strings.ToUpper(os.Args[1])

	for {
		// Print acts as a prompt
		fmt.Print("What is the current temperature in " + originUnit + " ? ")
		// Scanln accepts a variable for each argument accessed with the address of symbol
		_, err := fmt.Scanln(&originValue)

		if err != nil {
			printError(errReadingInput)
		}
		// Check if the input is for celsius
		if originUnit == "C" {
			convertToFahrenheit(originValue)
		} else {
			convertToCelsius(originValue)
		}
		// Print acts as a prompt - run this after the if/else statement executes
		fmt.Print("Would you like to convert another temperature ? (y/n) ")
		// Scanln accepts a variable for each argument accessed with the address of symbol
		_, e := fmt.Scanln(&shouldConvertAgain)

		if e != nil {
			printError(errReadingInput)
		}
		// Santize input, input will be uppercased and trim any whitespace
		if strings.ToUpper(strings.TrimSpace(shouldConvertAgain)) != "Y" {
			fmt.Println("Good bye!")
			break
		}
	}
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func convertToCelsius(value float64) {
	convertedValue := (value - 32) * 5 / 9
	fmt.Printf("%v F = %.0f C\n", value, convertedValue)
}

func convertToFahrenheit(value float64) {
	convertedValue := (value * 9 / 5) + 32
	fmt.Printf("%v C = %.0f F\n", value, convertedValue)
}
