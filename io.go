package core

import (
	"bufio"
	"io"
	"os"
)

// ForEachLine reads lines from the provided io.Reader and calls the provided
// function for each line read.
func ForEachLine(src io.Reader, fn func(string)) {
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		fn(scanner.Text())
	}
}

// GetLinesFromFile reads all lines from a file and returns them as a slice of strings.
func GetLinesFromFile(filePath string) ([]string, error) {
	var lines []string
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	ForEachLine(file, func(line string) {
		lines = append(lines, line)
	})
	return lines, nil
}
