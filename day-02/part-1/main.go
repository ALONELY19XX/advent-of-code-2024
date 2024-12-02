package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/ALONELY19XX/advent-of-code-2024/helpers"
)

const (
	MIN_DIFF = 1
	MAX_DIFF = 3
)

func main() {
	var filepath string
	helpers.SetFilepathFlag(&filepath)
	flag.Parse()

	lines, err := helpers.ReadInput(filepath)

	if err != nil {
		log.Fatal(err)
	}

	var totalSafeReports int

	// iterate over each report
	for _, line := range lines {
		levelsRaw := strings.Fields(line)

		// map raw levels (string) to ints
		levels := helpers.Map(levelsRaw, func(el string) int {
			num, err := strconv.Atoi(el)
			if err != nil {
				log.Fatal(err)
			}
			return num
		})

		var currLvl int
		var nextLvl int
		var isAsc bool

		// iterate levels within report
		for idx := range len(levels) - 1 {
			currLvl = levels[idx]
			nextLvl = levels[idx+1]

			diff := int(math.Abs(float64(currLvl) - float64(nextLvl)))

			// first iteration (first 2 levels) determine order (asc vs desc)
			if idx == 0 {
				if currLvl < nextLvl {
					isAsc = true
				} else if currLvl > nextLvl {
					isAsc = false
				} else {
					break
				}
			}

			if (isAsc && currLvl < nextLvl) || (!isAsc && currLvl > nextLvl) {
				if MIN_DIFF <= diff && diff <= MAX_DIFF {
					// everything fine so far.
				} else {
					break
				}
			} else {
				break
			}

			// if we reach this after comparing the last 2 levels, then its a safe report
			if idx == len(levels)-2 {
				totalSafeReports++
			}
		}
	}

	fmt.Println("Solution:", totalSafeReports) // Answer: 299
}
