package src

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Function or Global Variable or type variable if defined with a Capital letter: They will be available/accessible from outside the package else they are available only inside package
func ExtractCSVFiles(zipFile string, destDir string) {
	fmt.Println("Processing started...")
	reader, err := zip.OpenReader(zipFile)

	if err != nil {
		Log.Error(
			"Failed to open zip file",
			"zip_file", zipFile,
			"error", err,
		)
	}
	defer reader.Close()

	// Ensure destination directory exists. os.MkdirAll() creates the directory & if the directory already exists then no error
	// os.ModePerm is file permission 0777 (owner/group/others: read, write, execute)
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		Log.Error(
			"Failed to create destination directory",
			"destination_directory", destDir,
			"error", err,
		)
	}

	// Iterate through files in the zip
	for _, file := range reader.File {
		if !strings.HasSuffix(strings.ToLower(file.Name), ".csv") {
			continue
		}

		fmt.Println(file.Name, " of size ", file.FileInfo().Size(), "KB")

		readerCloser, err := file.Open()

		if err != nil {
			Log.Error(
				"Unable to open file inside zip",
				"File Name", file.Name,
				"error", err,
			)
		}

		// Take only the filename, ignore any folders from the ZIP, and safely place it inside the destination directory
		destPath := filepath.Join(destDir, filepath.Base(file.Name))

		outFile, err := os.Create(destPath)
		if err != nil {
			Log.Error(
				"Unable to create file at destination directory",
				"File Name", file.Name,
				"error", err,
			)
		}
		defer outFile.Close()

		// Copy contents
		_, err = io.Copy(outFile, readerCloser)
		if err != nil {
			Log.Error(
				"Unable to copy contents of .csv file",
				"File Name", file.Name,
				"error", err,
			)
		}
	}
	fmt.Println("Extracted files at location: ", destDir)
}
