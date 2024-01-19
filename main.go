package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/m/v2/wordcounter"
)

func main() {
	flag, filePath := os.Args[1], os.Args[2]
	fileParser, err := wordcounter.NewFileParser(filePath)
	if err != nil {
		fmt.Println(errors.New(err.Error()))
	}
	defer fileParser.File.Close()
	switch flag {
	case "-l":
		fileParser.LongestLine()
	case "-c":
		fileParser.TotalCharacters(true)
	case "-w":
		fileParser.TotalWords()
	default:
		fmt.Println("incorrect input")
	}
}
