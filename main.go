package main

import (
	"file-operation-using-go/src"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	zipPath := "test.zip"
	destPath := "extracted"

	src.ExtractCSVFiles(zipPath, destPath)
}
