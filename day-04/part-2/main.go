package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/ALONELY19XX/advent-of-code-2024/helpers"
)

func main() {
	var filepath string
	helpers.SetFilepathFlag(&filepath)
	flag.Parse()

	data, err := helpers.ReadInput(filepath)

	if err != nil {
		log.Fatal(err)
	}

	lineSep := "\r\n"
	lineLen := strings.Index(data, lineSep)
	// text as continuos block (newlines newlines/carriage-ret removed)
	text := strings.ReplaceAll(data, "\r\n", "")
	linesTotal := len(text) / lineLen

	totalXmas := 0

	for idx, ch := range text {
		// only check for 'A' since it always must be in center of X
		if ch == 'A' {
			currLine := idx / lineLen

			// check if there is enoush space around the A (at least 1 in each direction)
			isPaddingLeft := idx > (currLine * lineLen)
			isPaddingRight := idx < ((lineLen - 1) + (currLine * lineLen))
			isPaddingTop := currLine > 0
			isPaddingBottom := currLine < (linesTotal - 1)

			if isPaddingTop && isPaddingRight && isPaddingBottom && isPaddingLeft {
				// grab top-left, top-right, bottom-left and bottom-right neighbours
				tl := text[idx-lineLen-1]
				tr := text[idx-lineLen+1]
				bl := text[idx+lineLen-1]
				br := text[idx+lineLen+1]

				isDiagMas1 := false
				isDiagMas2 := false

				// opposite letters must be 'M' and 'A' pairs

				if (tl == byte('M') && br == byte('S')) || (tl == byte('S') && br == byte('M')) {
					isDiagMas1 = true
				}

				if (bl == byte('M') && tr == byte('S')) || (bl == byte('S') && tr == byte('M')) {
					isDiagMas2 = true
				}

				if isDiagMas1 && isDiagMas2 {
					totalXmas++
				}
			}
		}
	}

	fmt.Println("Solution:", totalXmas)
}
