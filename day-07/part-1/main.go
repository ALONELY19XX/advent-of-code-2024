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
	OP_PLUS = '+'
	OP_MULT = '*'
	OP_NONE = ' '
)

type Node struct {
	value  int
	result int
	left   *Node
	right  *Node
	op     byte
}

func main() {
	var filepath string
	helpers.SetFilepathFlag(&filepath)
	flag.Parse()

	lines, err := helpers.ReadInputLines(filepath)

	if err != nil {
		log.Fatal(err)
	}

	sum := 0

	for _, line := range lines {
		calibration := strings.Split(line, ":")

		result, _ := strconv.Atoi(calibration[0])
		operands := strings.Fields(calibration[1])

		tree := buildTree(operands, 0, OP_NONE, 0)

		if hasSolutionPath(tree, result) {
			sum += result
		}
	}

	fmt.Println("Solution:", sum)
}

func newNode(value int, op byte) *Node {
	return &Node{
		value:  value,
		op:     op,
		left:   nil,
		right:  nil,
		result: value,
	}
}

// construct binary tree where each operand is a node.
// for a given node:
// * the left child represents the addition with the next operand
// * the right child represents the multiplication with the next operand
// the first operand acts as the root of the tree
func buildTree(operands []string, idx int, op byte, result int) *Node {
	var node *Node

	// parse operand and add node
	if idx < len(operands) {
		operand, _ := strconv.Atoi(operands[idx])
		node = newNode(operand, op)
		if op == OP_PLUS {
			node.result = node.value + result
		} else if op == OP_MULT {
			node.result = node.value * result
		}
	}

	// keep adding nodes for all non-leaf nodes (last operand)
	if idx < len(operands)-1 {
		node.left = buildTree(operands, idx+1, OP_PLUS, node.result)
		node.right = buildTree(operands, idx+1, OP_MULT, node.result)
	}

	return node
}

// function which tests if a solution path (from root to leaf) exists for target value
func hasSolutionPath(node *Node, targetValue int) bool {
	// if we reached leaf, check if the sum matches target value
	if node.left == nil && node.right == nil {
		if node.result == targetValue {
			return true
		} else {
			return false
		}
	}

	// if not leaf, keep traversing the tree
	return hasSolutionPath(node.left, targetValue) || hasSolutionPath(node.right, targetValue)
}
