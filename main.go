package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	phonePattern := regexp.MustCompile(`\+90 \d{3} \d{3} \d{2} \d{2}`)
	var numbers []string

	file, err := os.Open("whatsapp.html")
	if err != nil {
		fmt.Println("Failed to open input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 10*1024*1024), 10*1024*1024) // 10MB buffer

	for scanner.Scan() {
		line := scanner.Text()
		matches := phonePattern.FindAllString(line, -1)
		numbers = append(numbers, matches...)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read input file:", err)
		return
	}

	if len(numbers) == 0 {
		fmt.Println("No phone numbers found.")
		return
	}

	tempFile := "temp_numbers.txt"
	if err := os.WriteFile(tempFile, []byte(strings.Join(numbers, "\n")), 0644); err != nil {
		fmt.Println("Failed to write temp file:", err)
		return
	}

	if err := removeDuplicates(tempFile, "numbers.txt"); err != nil {
		fmt.Println("Failed to remove duplicates:", err)
		return
	}

	_ = os.Remove(tempFile)
	fmt.Println("Unique phone numbers saved to numbers.txt")
}

func removeDuplicates(inputPath, outputPath string) error {
	inFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("unable to open input file: %v", err)
	}
	defer inFile.Close()

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("unable to create output file: %v", err)
	}
	defer outFile.Close()

	seen := make(map[string]bool)
	writer := bufio.NewWriter(outFile)
	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		line := scanner.Text()
		if !seen[line] {
			seen[line] = true
			if _, err := writer.WriteString(line + "\n"); err != nil {
				return fmt.Errorf("error writing to output file: %v", err)
			}
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("error flushing output file: %v", err)
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input file: %v", err)
	}

	return nil
}
