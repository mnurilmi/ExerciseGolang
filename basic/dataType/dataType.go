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
