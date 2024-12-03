package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/ALONELY19XX/advent-of-code-2024/helpers"
)

func main() {
	var filepath string
	helpers.SetFilepathFlag(&filepath)
	flag.Parse()

	mem, err := helpers.ReadInput(filepath)

	if err != nil {
		log.Fatal(err)
	}

	// matches "do()", "don't()" and/or "mul(<num1>, <num2>)" (where num1/num2 are 3-digit numbers)
	mulRegex := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`)

	// get all matches, where match consists of 3 values.
	// Consider the following matching pattern "xyz**mul(1, 23)**xyz"
	// * 1st value: entire matched string, e.g. "mul(1, 23)"
	// * 2nd value: first submatch "1"
	// * 3rd value: second submatch "23"
	// -- OR --
	//  where match consists of 2 values.
	// * 1st value: entire matched string, e.g. "do()" or "don't()"
	// * 2nd value: empty string "" (since no submatch pattern was matched for these cases)
	matches := mulRegex.FindAllStringSubmatch(mem, -1)

	res := 0

	isEnabled := true

	for _, match := range matches {
		// get matched operation: "do()", "don't()" or "mul(<num1>, <num2>)"
		op := match[0]

		if op == "do()" {
			isEnabled = true
		} else if op == "don't()" {
			isEnabled = false
		} else {
			if isEnabled {
				firstOp, _ := strconv.Atoi(match[1])
				secondOp, _ := strconv.Atoi(match[2])
				res += firstOp * secondOp
			}
		}
	}

	fmt.Println("Solution:", res)
}
