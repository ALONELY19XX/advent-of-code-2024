package main

import (
	"flag"
	"fmt"
	"log"
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

	// left list to store left location IDs
	llist := make([]int64, len(lines))
	// occurrence map to store right location ID occurrences
	occurences := make(map[int64]int64)

	for idx, line := range lines {
		locIds := strings.Fields(line)

		// parse left location iD and add it to the left list
		locId, err := strconv.ParseInt(locIds[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		llist[idx] = locId

		// parse right location ID and register it in occurrence map
		locId, err = strconv.ParseInt(locIds[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		_, ok := occurences[locId]
		if ok {
			occurences[locId]++
		} else {
			occurences[locId] = 1
		}

	}

	var similarityScore int64

	// accumulate total similarity score based on left location ID's occurrence count
	for _, locId := range llist {
		if occ, ok := occurences[locId]; ok {
			similarityScore += locId * occ
		}

	}

	fmt.Println("Solution:", similarityScore) // Answer: 21070419
}
