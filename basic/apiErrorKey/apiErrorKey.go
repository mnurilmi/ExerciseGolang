package main

import (
	"fmt"
	"strconv"
)

var errorCodes = map[int]string{
	0: "No Error",
	1: "Incorrect input",
	2: "The server encounters internal error",
	4: "The server is overloaded by too much traffic",
	8: "You are not authorized to proceed with the input",
}

func main() {
	var in string
	var inInt int

	if _, err := fmt.Print("input:"); err != nil {
		fmt.Println(err.Error())
	}

	if _, err := fmt.Scanf("%s", &in); err != nil {
		fmt.Println(err.Error())
	}

	output := []string{}
	inInt, err := strconv.Atoi(in)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if inInt == 0 {
		output = append(output, errorCodes[0])
	} else {
		for inInt > 0 {
			if inInt >= 8 {
				output = append(output, errorCodes[8])
				inInt -= 8
			} else if inInt >= 4 {
				output = append(output, errorCodes[4])
				inInt -= 4
			} else if inInt >= 2 {
				output = append(output, errorCodes[2])
				inInt -= 2
			} else if inInt >= 1 {
				output = append(output, errorCodes[1])
				inInt -= 1
			}
		}
	}

	if _, err := fmt.Println("Output:", output); err != nil {
		fmt.Println(err.Error())
	}
}

/*
Imagine you are a software developer who has been tasked with integrating a third-party API into your company's system. This API is crucial for your project, but it has a unique way of communicating errors. Instead of conventional error messages, it sends integer values, each representing a specific error.

You've been given the following mappings of error codes to their meanings:

Code	Description
1	Incorrect input
2	The server encounters internal error
4	The server is overloaded by too much traffic
8	You are not authorized to proceed with the input
If there is no errors, then your function should return ["No Error"]

To complicate matters, the API might combine error codes to represent multiple problems occurring at the same time.

The output should be an array (or a slice in Go) of error messages.

Examples
Input: 0
Output: ["No Error"]

Input: 8
Output: ["You are not authorized to proceed with the input"]

Input: 3
Output: ["Incorrect input", "The server encounters internal error"]
*/
