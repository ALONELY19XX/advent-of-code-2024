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

		badLvl := false

		levelsRaw := strings.Fields(line)

		// map raw levels (string) to ints
		levels := helpers.Map(levelsRaw, func(el string) int {
			num, err := strconv.Atoi(el)
			if err != nil {
				log.Fatal(err)
			}
			return num
		})

		// flag to determine order of levels (asc vs desc)
		isAsc := false

		for idx := range len(levels) - 1 {
			curr := levels[idx]
			next := levels[idx+1]
			diff := curr - next
			diffAbs := int(math.Abs(float64(curr) - float64(next)))

			// determine order on first pair (or abort)
			if idx == 0 {
				if diff < 0 {
					isAsc = true
				} else if diff > 0 {
					isAsc = false
				} else {
					badLvl = true
					break
				}
			}

			// if pair violates against determined order: abort
			if (isAsc && diff > 0) || (!isAsc && diff < 0) {
				badLvl = true
				break
			}

			// check if fluctuation is within allowed limits
			if MIN_DIFF <= diffAbs && diffAbs <= MAX_DIFF {
				continue
			} else {
				badLvl = true
				break
			}
		}

		if badLvl {
			continue
		} else {
			totalSafeReports++
		}
	}

	fmt.Println("Solution:", totalSafeReports) // Answer: 299
}
