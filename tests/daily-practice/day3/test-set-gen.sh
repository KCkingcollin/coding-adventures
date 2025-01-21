#!/bin/bash

testSetFile="test-set"
answerSetFile="answer-set"

touch "$testSetFile"
touch "$answerSetFile"

for _ in {1..50}; do
    n=$(((RANDOM % 1000)+1))
    input="$n"$'\n'
    for ((i=0; i<n; i++)); do
        for ((j=0; j<3; j++)); do
            input+="$((RANDOM % 2))"
            if [ "$j" -gt 1 ]; then
                input+=$'\n'
            else
                input+=" "
            fi
        done
    done
    
    output=$(echo "$input" | go run ../main.go)
    
    input="{\""$'\n'"$input"$'\n'"\"}"
    output="{\""$'\n'"$output"$'\n'"\"}"
    echo "$input" >> "$testSetFile"
    echo "$output" >> "$answerSetFile"
done

echo "Files written successfully"
