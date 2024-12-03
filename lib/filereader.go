package lib

import (
	"bufio"
	"os"
)

func ReadFile(filepath string) ([]string, error) {
	readFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines, nil
}
