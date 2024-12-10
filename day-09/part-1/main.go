package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ALONELY19XX/advent-of-code-2024/helpers"
)

const EMPTY_BLOCK_ID = -1

func main() {
	var filepath string
	helpers.SetFilepathFlag(&filepath)
	flag.Parse()

	diskmap, err := helpers.ReadInput(filepath)

	if err != nil {
		log.Fatal(err)
	}

	// list to keep track of the arrangement of empty spaces (denoted with '-1') and fileId blocks
	var blocks []int
	// list which keeps track of all free space indices
	var freeSpaces []int

	// alternate between numbers and populate the blocks list
	currFileId := 0
	for idx, block := range diskmap {
		isFileBlock := idx%2 == 0
		num := int(block) - '0'
		if isFileBlock {
			for range num {
				blocks = append(blocks, currFileId)
			}
			currFileId++
		} else {
			for range num {
				// if free space, add index to tracking list
				freeSpaces = append(freeSpaces, len(blocks))
				blocks = append(blocks, EMPTY_BLOCK_ID)
			}
		}
	}

	// from right to left: change non-free block with frontmost free space
	freeSpacesIdx := 0
	for i := len(blocks) - 1; i >= 0; i-- {
		if freeSpacesIdx < len(freeSpaces) {
			freeSpace := freeSpaces[freeSpacesIdx]
			if blocks[i] >= 0 && i > freeSpace {
				blocks[freeSpace] = blocks[i]
				blocks[i] = EMPTY_BLOCK_ID
				freeSpacesIdx++
			}
		} else {
			break
		}
	}

	checksum := int64(0)

	// calculate checksum for all fileIds (left to right) until we encounter free space
	for idx, fileId := range blocks {
		if fileId != EMPTY_BLOCK_ID {
			checksum += int64(idx * fileId)
		} else {
			break
		}
	}

	fmt.Println("Solution:", checksum)
}
