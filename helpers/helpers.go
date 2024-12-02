package helpers

import (
	"bufio"
	"flag"
	"os"
)

func ReadInput(filepath string) ([]string, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func SetFilepathFlag(arg *string) {
	flag.StringVar(arg, "filepath", ".", "The path to the challenge's input file.")
}

func Map[T, K any](input []T, fn func(T) K) []K {
	output := make([]K, len(input))
	for i, t := range input {
		output[i] = fn(t)
	}
	return output
}
