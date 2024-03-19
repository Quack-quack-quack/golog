package parser

import (
	"bufio"
	"fmt"
	"os"
)

func OpenFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", filePath)
	}
	return file, nil
}

func CloseFile(file *os.File) {
	if file != nil {
		file.Close()
	}
}

func ReadLines(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
