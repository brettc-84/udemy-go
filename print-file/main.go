package main

import (
	"fmt"
	"io"
	"os"
)

type filePrinter struct{}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Filename must be provided")
		os.Exit(1)
	}

	file, _ := os.Open(os.Args[1])
	io.Copy(filePrinter{}, file)
}

func (filePrinter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
