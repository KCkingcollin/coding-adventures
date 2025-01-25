package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run <testSet_location> <answerSet_location> <command>")
		os.Exit(1)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v\n", err)
	}
	fmt.Printf("Working DIR: %s\n", dir)

	testSet := os.Args[1]
	answerSet := os.Args[2]
	command := os.Args[3]

	// Replace "." with the current working directory
	command = strings.Replace(command, "./", dir+"/", 1)

	// Split the command into command and arguments
	args := strings.Split(command, " ")

	// Check if the command exists in the PATH
	if _, err := exec.LookPath(args[0]); err != nil {
		fmt.Printf("Error: Command '%s' does not exist.\n", args[0])
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
		cmd := exec.Command(args[0], args[1:]...) // Pass command and its arguments
		cmd.Stdin = strings.NewReader(inputBlock)
		actualOutputBytes, err := cmd.CombinedOutput() // Capture both stdout and stderr
		if err != nil {
			fmt.Printf("Error running command: %v\n", err)
			fmt.Printf("Stderr: %s\n", string(actualOutputBytes))
			os.Exit(1)
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
