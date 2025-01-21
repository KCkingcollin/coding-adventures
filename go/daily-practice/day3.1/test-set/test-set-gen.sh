#!/bin/bash

testSetFile="test-set"
answerSetFile="answer-set"

touch "$testSetFile"  # Create or truncate the file
touch "$answerSetFile" # Create or truncate the file

for _ in {1..50}; do
    n=$((RANDOM % 150 + 1))
    input=""
    input+="$n"$'\n'
    for ((i=0; i<n; i++)); do
        leftOrRight=$((RANDOM % 2))
        plusOrMinus=$((RANDOM % 2))
        if ((plusOrMinus == 0)); then
            op="++"
        else
            op="--"
        fi
        if ((leftOrRight == 0)); then
            value1="$op"
            value2="X"
        else
            value1="X"
            value2="$op"
        fi
        input+="$value1$value2"$'\n'
    done
    
    output=$(echo "$input" | go run ../main.go)
    
    input="{\""$'\n'"$input"$'\n'"\"}"
    output="{\""$'\n'"$output"$'\n'"\"}"
    echo "$input" >> "$testSetFile"
    echo "$output" >> "$answerSetFile"
done

echo "Files written successfully"
