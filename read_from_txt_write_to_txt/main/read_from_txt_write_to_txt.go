package main

import (
	"bufio"
	"fmt"
	"os"
)

// copyFi;e copies content from the source file to the destination file
// It reads/writes line by line
func copyFile(src, dest string) (int, error) {
	// Open the source file
	inputFile, err := os.Open(src)
	if err != nil {
		return 0, fmt.Errorf("Error opening input file : %w", err)
	}
	defer inputFile.Close()

	// Create the destination file
	outputFile, err := os.Create(dest)
	if err != nil {
		return 0, fmt.Errorf("Error creating output file: %w", err)
	}
	defer outputFile.Close()

	// Initiate buffered reader & writer
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	totalBytesWriten := 0

	// Read & write each line from sorce file to destination file
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			return totalBytesWriten, fmt.Errorf("Error reading inpur file : %w", err)
		}
		n, err := writer.WriteString(line + "\n")
		if err != nil {
			return totalBytesWriten, fmt.Errorf("Error writing to output file: %w", err)
		}
		totalBytesWriten += n
	}
	// Flush the buffered writer
	err = writer.Flush()
	if err != nil {
		return totalBytesWriten, fmt.Errorf("Error flushing writer: %w", err)
	}
	return totalBytesWriten, nil
}

func main() {
	srcFile := "testSrc.txt"
	destFile := "testDest.txt"

	totalBytes, err := copyFile(srcFile, destFile)
	if err != nil {
		fmt.Errorf("File copy failed! ", err)
	} else {
		fmt.Errorf("File copied successfully! total bytes writen: %d\n", totalBytes)
	}
}
