package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/ALONELY19XX/advent-of-code-2024/helpers"
)

const (
	MOVE_UP = iota
	MOVE_RIGHT
	MOVE_DOWN
	MOVE_LEFT
	MOVE_END
)

const OBSTACLE_MARK = '#'

func main() {
	var filepath string
	helpers.SetFilepathFlag(&filepath)
	flag.Parse()

	data, err := helpers.ReadInput(filepath)

	if err != nil {
		log.Fatal(err)
	}

	cols := strings.Index(data, "\r\n")
	// remove newlines so the entire grid is one continuos string of tiles
	grid := strings.ReplaceAll(data, "\r\n", "")

	guardPos := strings.Index(grid, "^")
	moveDir := MOVE_UP

	// map to track unique tiles which were visited
	visits := make(map[int]bool)
	visits[guardPos] = true

	for {
		switch moveDir {
		case MOVE_UP:
			newPosIdx, newMoveDir := tryMovingUp(grid, cols, guardPos)
			guardPos = newPosIdx
			moveDir = newMoveDir
		case MOVE_RIGHT:
			newPosIdx, newMoveDir := tryMovingRight(grid, cols, guardPos)
			guardPos = newPosIdx
			moveDir = newMoveDir
		case MOVE_DOWN:
			newPosIdx, newMoveDir := tryMovingDown(grid, cols, guardPos)
			guardPos = newPosIdx
			moveDir = newMoveDir
		case MOVE_LEFT:
			newPosIdx, newMoveDir := tryMovingLeft(grid, cols, guardPos)
			guardPos = newPosIdx
			moveDir = newMoveDir
		}

		if moveDir == MOVE_END {
			break
		}

		visits[guardPos] = true
	}

	distinctVisits := len(visits)

	fmt.Println("Solution:", distinctVisits)
}

func tryMovingUp(grid string, cols int, posIdx int) (newPosIdx int, newMoveDir int) {
	// new index if we move one tile up
	frontIdx := posIdx - cols

	// if we didn't escape the grid
	if frontIdx >= 0 {
		if grid[frontIdx] != OBSTACLE_MARK {
			return frontIdx, MOVE_UP
		} else {
			return posIdx, MOVE_RIGHT
		}
	}
	// we escaped grid through the top boundary
	return posIdx, MOVE_END
}

func tryMovingRight(grid string, cols int, posIdx int) (newPosIdx int, newMoveDir int) {
	currRow := (posIdx / cols)
	// new index if we move one tile right
	frontIdx := posIdx + 1
	newRow := (frontIdx / cols)

	// if we didn't escape the grid
	if currRow == newRow && frontIdx < len(grid) {
		if grid[frontIdx] != OBSTACLE_MARK {
			return frontIdx, MOVE_RIGHT
		} else {
			return posIdx, MOVE_DOWN
		}
	}
	// we escaped grid through the right boundary
	return posIdx, MOVE_END
}

func tryMovingDown(grid string, cols int, posIdx int) (newPosIdx int, newMoveDir int) {
	// new index if we move one tile down
	frontIdx := posIdx + cols

	// if we didn't escape the grid
	if frontIdx < len(grid) {
		if grid[frontIdx] != OBSTACLE_MARK {
			return frontIdx, MOVE_DOWN
		} else {
			return posIdx, MOVE_LEFT
		}
	}
	// we escaped grid through the bottom boundary
	return posIdx, MOVE_END
}

func tryMovingLeft(grid string, cols int, posIdx int) (newPosIdx int, newMoveDir int) {
	currRow := (posIdx / cols)
	// new index if we move one tile left
	frontIdx := posIdx - 1
	newRow := (frontIdx / cols)

	// if we didn't escape the grid
	if currRow == newRow && frontIdx >= 0 {
		if grid[frontIdx] != OBSTACLE_MARK {
			return frontIdx, MOVE_LEFT
		} else {
			return posIdx, MOVE_UP
		}
	}
	// we escaped grid through the left boundary
	return posIdx, MOVE_END
}
