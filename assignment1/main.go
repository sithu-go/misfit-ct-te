package main

import (
	"fmt"
	"unicode"
)

func main() {
	testCases := []struct {
		input  string
		output string
	}{
		{"Var-----___1_int", "1"},
		{"Q2q-q", "2"},
		{"eef3243**s", "3"},
	}

	for _, testCase := range testCases {
		result := grabFirstNumeric(testCase.input)
		if result == testCase.output {
			fmt.Printf("Input: %s, Expected: %s, Result: %s (Pass)\n", testCase.input, testCase.output, result)
		} else {
			fmt.Printf("Input: %s, Expected: %s, Result: %s (Fail)\n", testCase.input, testCase.output, result)
		}
	}

	// for grabbing all numbers
	// for _, testCase := range testCases {
	// 	result := grabAllNumbers(testCase.input)
	// 	fmt.Printf("Result: %v\n", result)
	// }
}

func grabFirstNumeric(input string) string {
	for _, char := range input {
		if unicode.IsDigit(char) {
			return string(char)
		}
	}
	return ""
}

func grabAllNumbers(input string) []string {
	var numbers []string
	currentNumber := ""

	for _, char := range input {
		if unicode.IsDigit(char) {
			currentNumber += string(char)
		} else if currentNumber != "" {
			fmt.Println(currentNumber)
			numbers = append(numbers, currentNumber)
			currentNumber = ""
		}
	}

	if currentNumber != "" {
		numbers = append(numbers, currentNumber)
	}

	return numbers
}
