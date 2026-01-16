package src

import (
	"archive/zip"
	"fmt"
	"strings"
)

// Function or Global Variable or type variable if defined with a Capital letter: They will be available/accessible from outside the package else they are available only inside package
func ExtractCSVFiles(sourcePath string, destPath string) {
	fmt.Println("Processing started...")
	reader, err := zip.OpenReader(sourcePath)

	if err != nil {
		Log.Error(
			"Failed to open zip file",
			"zip_file", sourcePath,
			"error", err,
		)
	}

	for _, file := range reader.File {
		if !strings.HasSuffix(strings.ToLower(file.Name), ".csv") {
			continue
		}

		fmt.Println(file.Name, " of size ", file.FileInfo().Size(), "KB")
	}

}
