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

	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// get all matches, where match consists of 3 values.
	// Consider the following matching pattern "xyz**mul(1, 23)**xyz"
	// * 1st value: entire matched string, e.g. "mul(1, 23)"
	// * 2nd value: first submatch "1"
	// * 3rd value: second submatch "23"
	matches := mulRegex.FindAllStringSubmatch(mem, -1)

	res := 0

	for _, match := range matches {
		firstOp, _ := strconv.Atoi(match[1])
		secondOp, _ := strconv.Atoi(match[2])
		res += firstOp * secondOp
	}

	fmt.Println("Solution:", res)
}
