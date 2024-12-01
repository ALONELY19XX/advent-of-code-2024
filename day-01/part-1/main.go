package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/ALONELY19XX/advent-of-code-2024/helpers"
)

func main() {
	var filepath string
	helpers.ParseFilepath(&filepath)
	flag.Parse()

	lines, err := helpers.ReadInput(filepath)

	if err != nil {
		log.Fatal(err)
	}

	// left and right lists to store corresponding location IDs
	llist := make([]int64, len(lines))
	rlist := make([]int64, len(lines))

	for idx, line := range lines {
		locIds := strings.Fields(line)

		// parse left location ID and add it to the left list
		locId, err := strconv.ParseInt(locIds[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		llist[idx] = locId

		// parse right location ID and add it to the right list
		locId, err = strconv.ParseInt(locIds[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		rlist[idx] = locId
	}

	// Sort both lists in ascending order
	slices.Sort(llist)
	slices.Sort(rlist)

	var dist int64

	// accumulate absolute differences between ordered location ID pairs
	for idx := range len(lines) {
		diff := float64(llist[idx] - rlist[idx])
		dist += int64(math.Abs(diff))
	}

	fmt.Println("Solution:", dist) // Answer: 1223326
}
