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

	lineSep := "\r\n"
	lineLen := strings.Index(data, lineSep)
	// text as continuos block (newlines newlines/carriage-ret removed)
	text := strings.ReplaceAll(data, "\r\n", "")
	linesTotal := len(text) / lineLen

	totalXmas := 0

	for idx, ch := range text {
		if ch == 'X' {
			currLine := idx / lineLen
			// scan right
			if idx < (lineLen-3)+(currLine*lineLen) {
				ch2 := text[idx+1]
				ch3 := text[idx+2]
				ch4 := text[idx+3]
				composeAndCheckWord(&totalXmas, byte(ch), ch2, ch3, ch4)
			}
			// scan left
			if idx > 2+(currLine*lineLen) {
				ch2 := text[idx-1]
				ch3 := text[idx-2]
				ch4 := text[idx-3]
				composeAndCheckWord(&totalXmas, byte(ch), ch2, ch3, ch4)
			}
			// scan top
			if currLine > 2 {
				ch2 := text[idx-lineLen]
				ch3 := text[idx-lineLen*2]
				ch4 := text[idx-lineLen*3]
				composeAndCheckWord(&totalXmas, byte(ch), ch2, ch3, ch4)
			}
			// scan bottom
			if currLine < (linesTotal - 3) {
				ch2 := text[idx+lineLen]
				ch3 := text[idx+lineLen*2]
				ch4 := text[idx+lineLen*3]
				composeAndCheckWord(&totalXmas, byte(ch), ch2, ch3, ch4)
			}
			// scan top-left
			if (idx > 2+(currLine*lineLen)) && (currLine > 2) {
				ch2 := text[(idx-lineLen)-1]
				ch3 := text[(idx-lineLen*2)-2]
				ch4 := text[(idx-lineLen*3)-3]
				composeAndCheckWord(&totalXmas, byte(ch), ch2, ch3, ch4)
			}
			// scan top-right
			if (idx < (lineLen-3)+(currLine*lineLen)) && (currLine > 2) {
				ch2 := text[(idx-lineLen)+1]
				ch3 := text[(idx-lineLen*2)+2]
				ch4 := text[(idx-lineLen*3)+3]
				composeAndCheckWord(&totalXmas, byte(ch), ch2, ch3, ch4)
			}
			// scan bottom-left
			if (currLine < (linesTotal - 3)) && (idx > 2+(currLine*lineLen)) {
				ch2 := text[(idx+lineLen)-1]
				ch3 := text[(idx+lineLen*2)-2]
				ch4 := text[(idx+lineLen*3)-3]
				composeAndCheckWord(&totalXmas, byte(ch), ch2, ch3, ch4)
			}
			// scan bottom-right
			if (currLine < (linesTotal - 3)) && (idx < (lineLen-3)+(currLine*lineLen)) {
				ch2 := text[(idx+lineLen)+1]
				ch3 := text[(idx+lineLen*2)+2]
				ch4 := text[(idx+lineLen*3)+3]
				composeAndCheckWord(&totalXmas, byte(ch), ch2, ch3, ch4)
			}
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Solution:", totalXmas)
}

func composeAndCheckWord(counter *int, bytes ...byte) {
	if string(bytes) == "XMAS" {
		(*counter)++
	}
}
