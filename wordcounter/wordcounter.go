package wordcounter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FileParser struct {
	File *os.File
}

func NewFileParser(filePath string) (FileParser, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return FileParser{}, err
	}

	return FileParser{File: file}, nil
}

func (f *FileParser) LongestLine() {
	length := 0
	scanner := bufio.NewScanner(f.File)
	for scanner.Scan() {
		lineLength := len(scanner.Text())
		if lineLength > length {
			length = lineLength
		}
	}

	fmt.Println("Length of the longest line is ", length, " .")
}

func (f *FileParser) TotalCharacters(characters bool) {
	length := 0
	scanner := bufio.NewScanner(f.File)
	for scanner.Scan() {
		length += len(scanner.Text())
	}
	if characters {
		fmt.Println("Total number of characters is ", length, ".")
	} else {
		fmt.Println("Total number of bytes is ", length, ".")
	}
}

func (f *FileParser) TotalLines() {
	lines := 0
	Scanner := bufio.NewScanner(f.File)
	for Scanner.Scan() {
		lines++
	}
	fmt.Println("Total number of lines is ", lines, ".")
}

func (f *FileParser) TotalWords() {
	words := 0
	Scanner := bufio.NewScanner(f.File)
	for Scanner.Scan() {
		words += len(strings.Fields(Scanner.Text()))
	}
	fmt.Println("total words in the text file are ", words, ".")
}
