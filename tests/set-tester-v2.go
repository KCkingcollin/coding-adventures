package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"os/exec"
	"log"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run <testSet_location> <answerSet_location> <command>")
		os.Exit(1)
	}

	testSet := os.Args[1]
	answerSet := os.Args[2]
	command := os.Args[3]

	// Check if files exist
	if _, err := os.Stat(testSet); os.IsNotExist(err) {
		fmt.Printf("Error: Test set file '%s' does not exist.\n", testSet)
		os.Exit(1)
	}

	if _, err := os.Stat(answerSet); os.IsNotExist(err) {
		fmt.Printf("Error: Answer set file '%s' does not exist.\n", answerSet)
		os.Exit(1)
	}

	// Open test and answer files
	testFile, err := os.Open(testSet)
	if err != nil {
		log.Fatal(err)
	}
	defer testFile.Close()

	answerFile, err := os.Open(answerSet)
	if err != nil {
		log.Fatal(err)
	}
	defer answerFile.Close()

	// Create bufio readers for both files
	testReader := bufio.NewReader(testFile)
	answerReader := bufio.NewReader(answerFile)

	lineNum := 1
	blockNum := 1

	for {
		// Read a line from both the test and answer files
		testLine, err := testReader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				log.Fatal(err)
			}
			break
		}

		expectedLine, err := answerReader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				log.Fatal(err)
			}
			break
		}

		inputBlock := testLine
		expectedOutputBlock := expectedLine

		// Read test block until both '{' and '}' are present
		for {
			testLine, err = testReader.ReadString('\n')
			if err != nil {
				break
			}
			inputBlock += testLine

			if strings.Contains(inputBlock, "{\"") && strings.Contains(inputBlock, "\"}") {
				break
			}
		}

		// Read expected output block until both '{' and '}' are present
		for {
			expectedLine, err = answerReader.ReadString('\n')
			if err != nil {
				break
			}
			expectedOutputBlock += expectedLine

			if strings.Contains(expectedOutputBlock, "{\"") && strings.Contains(expectedOutputBlock, "\"}") {
				break
			}
		}

		// Remove first and last line from input and expected output blocks
		inputBlock = strings.Join(strings.Split(inputBlock, "\n")[1:len(strings.Split(inputBlock, "\n"))-2], "\n") + "\n"
		expectedOutputBlock = strings.Join(strings.Split(expectedOutputBlock, "\n")[1:len(strings.Split(expectedOutputBlock, "\n"))-2], "\n") + "\n"

		// Count the number of lines in the input block
		numLines := len(strings.Split(inputBlock, "\n"))

		// Run the command and capture the output
		fmt.Printf("Running test number %d\n", blockNum)
		cmd := exec.Command(command)
		cmd.Stdin = strings.NewReader(inputBlock)
		actualOutputBytes, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		actualOutput := string(actualOutputBytes)

		// Compare the actual output to the expected output
		if actualOutput != expectedOutputBlock {
			fmt.Printf("Output for block number '%d' (starting at line number '%d') does not match the expected output.\n", blockNum, lineNum)
			fmt.Printf("Expected: '%s'\n", expectedOutputBlock)
			fmt.Printf("Got: '%s'\n", actualOutput)
			fmt.Printf("Input: '%s'\n", inputBlock)
			os.Exit(1)
		} else {
			fmt.Println("Passed")
		}

		lineNum += numLines
		blockNum++
	}

	fmt.Println("All outputs match")
}
