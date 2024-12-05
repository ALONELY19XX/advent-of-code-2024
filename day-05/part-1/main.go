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
	helpers.SetFilepathFlag(&filepath)
	flag.Parse()

	lines, err := helpers.ReadInputLines(filepath)

	if err != nil {
		log.Fatal(err)
	}

	// get index of empty line which acts as separator between order-rules and updates
	sepIdx := -1
	for idx, line := range lines {
		if len(line) == 0 {
			sepIdx = idx
			break
		}
	}

	rules := lines[:sepIdx]
	updates := lines[sepIdx+1:]

	// map between a page (key) and list of pages (value) which may only occur afterwards
	orderMap := make(map[int][]int)

	// populate order map
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		orderMap[l] = append(orderMap[l], r)
	}

	// list which will be populated with middle page of valid updates
	var middlePages []int

	// for every update sequence:
	for _, update := range updates {
		pages := strings.Split(update, ",")
		invalidOrder := false
		// map which maps page (key) to usage (value) represented as boolean.
		// the value "true" will indicate that the pages was seen already
		usageMap := make(map[int]bool)

		for _, page := range pages {
			num, _ := strconv.Atoi(page)

			// if current page in update sequence has a order-map entry:
			if list, ok := orderMap[num]; ok {
				// iterate through every page that may only occur afterwards
				for _, value := range list {
					if v, ok := usageMap[value]; ok {
						// if such a page was already seen we know that the order rule was broken
						if v {
							invalidOrder = true
							break
						}
					}
				}
			}

			if !invalidOrder {
				usageMap[num] = true
			} else {
				break
			}
		}

		// track middle page of valid update sequence
		if !invalidOrder {
			middlePage, _ := strconv.Atoi(pages[len(pages)/2])
			middlePages = append(middlePages, middlePage)
		}
	}

	sum := 0

	for _, page := range middlePages {
		sum += page
	}

	fmt.Println("Solution:", sum)
}
