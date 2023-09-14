package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	// Test cases
	testCases := []struct {
		input  string
		output bool
	}{
		{"172.316.254.1", false}, // Please note that: it should be false, cause 316 is larger than 255
		{"0.1.1.256", false},
		{"1.1.1.1a", false},
		{"1", false},
		{"64.233.16.00", false},
		{"7283728", false},
	}

	// check with standard built in library
	for _, testCase := range testCases {
		result := checkIPAddressWithBuiltInLibrary(testCase.input)
		if result == testCase.output {
			fmt.Printf("Input: \"%s\", Expected: %v, result: %v (Pass)\n", testCase.input, testCase.output, result)
		} else {
			fmt.Printf("Input: \"%s\", Expected: %v, result: %v (Fail)\n", testCase.input, testCase.output, result)
		}
	}

	// check manual
	// for _, testCase := range testCases {
	// 	if isIPAddressValid(testCase.input) == testCase.output {
	// 		fmt.Printf("Input: \"%s\", Expected: %v result: true (Pass)\n", testCase.input, testCase.output)
	// 	} else {
	// 		fmt.Printf("Input: \"%s\", Expected: %v result: false (Fail)\n", testCase.input, testCase.output)
	// 	}
	// }
}

func checkIPAddressWithBuiltInLibrary(ip string) bool {
	return net.ParseIP(ip) != nil
}

func isIPAddressValid(ipAddressStr string) bool {
	// Check if the input string is empty or too long
	if len(ipAddressStr) == 0 || len(ipAddressStr) > 15 {
		return false
	}

	// Split the IP address into segments using the dot as a delimiter
	segments := strings.Split(ipAddressStr, ".")

	// Check if there are exactly 4 segments
	if len(segments) != 4 {
		return false
	}

	for _, segment := range segments {
		// Try to parse the segment as an integer
		num, err := strconv.Atoi(segment)
		if err != nil {
			return false
		}

		// Check if the integer is within the valid range [0, 255]
		if num < 0 || num > 255 {
			return false
		}

		// Check for leading zeros in each segment
		if len(segment) > 1 && segment[0] == '0' {
			return false
		}
	}

	// If all checks passed,
	return true
}
