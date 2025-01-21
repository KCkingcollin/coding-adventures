#!/bin/bash

# Create the test and answer set files
testSetFile="test-set"
answerSetFile="answer-set"

touch "$testSetFile"  # Create or truncate the file
touch "$answerSetFile" # Create or truncate the file

# Generate 50 random numbers, run the command and store inputs and outputs
for _ in {1..50}; do
    # Generate a random test value
    testValue=$(((RANDOM % 100)+1))
    input="$testValue"
    
    # Run the command and capture the output
    output=$(echo "$input" | go run ../main.go)
    
    # Append the input and output to the respective files
    input="{\""$'\n'"$input"$'\n'"\"}"
    output="{\""$'\n'"$output"$'\n'"\"}"
    echo "$input" >> "$testSetFile"
    echo "$output" >> "$answerSetFile"
done

echo "Files written successfully"
