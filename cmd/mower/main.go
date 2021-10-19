package main

import (
	"fmt"
	"jellysmack/internal/pkg/parser"
	"os"
)

func usage(programName string) {
	fmt.Println("usage :\t\t", programName, "path/to/movements.txt")
	fmt.Println("example :\t", programName, "./tests/test.txt")
}

func main() {
	programName, args := os.Args[0], os.Args[1:]
	if len(args) != 1 {
		usage(programName)
		return
	}

	fileParser := parser.NewFileParser(args[0])
	b, err := fileParser.Extract()
	if err != nil {
		panic(err)
	}
	b.Play()
}
