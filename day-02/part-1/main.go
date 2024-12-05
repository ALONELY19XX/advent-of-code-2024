package main

import (
	"flag"
	"fmt"
	"log"
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

	lines, err := helpers.ReadInputLines(filepath)

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

		diffs := make([]int, len(levels)-1)

		// flag which tracks if a zero-diff is detected.
		// in this case we can abort early and dont need to check for
		// asc or desc ordering
		zeroDiffExists := false

		for idx := range len(levels) - 1 {
			curr := levels[idx]
			next := levels[idx+1]
			diff := curr - next
			if diff == 0 {
				zeroDiffExists = true
				break
			}
			diffs[idx] = diff
		}

		if zeroDiffExists {
			continue
		}

		if isAllValidAsc(diffs) || isAllValidDesc(diffs) {
			totalSafeReports++
		}

	}

	fmt.Println("Solution:", totalSafeReports)
}

// check if all levels are ascending and within limits
// a fully ascending levels report is represented by a fully negative diffs list
// e.g. consider report "1 3 4 7", then the diffs list will be:
// [(1-3), (3-4), (4-7)] = [-2, -1, -3]
func isAllValidAsc(diffs []int) bool {
	isAllValid := true
	for _, diff := range diffs {
		if !(-MAX_DIFF <= diff && diff <= -MIN_DIFF) {
			isAllValid = false
		}
	}
	return isAllValid
}

// check if all levels are descending and within limits
// a fully descending levels report is represented by a fully positive diffs list
// e.g. consider report "7 4 3 1", then the diffs list will be:
// [(7-4), (4-3), (3-1)] = [3, 1, 2]
func isAllValidDesc(diffs []int) bool {
	isAllValid := true
	for _, diff := range diffs {
		if !(MIN_DIFF <= diff && diff <= MAX_DIFF) {
			isAllValid = false
		}
	}
	return isAllValid
}
