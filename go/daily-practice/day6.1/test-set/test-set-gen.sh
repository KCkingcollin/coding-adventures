#!/bin/bash

testSetFile="test-set"
answerSetFile="answer-set"

touch "$testSetFile"  
touch "$answerSetFile"

for _ in {1..50}; do
    input=""
    n=$((RANDOM % 100 + 1))
    for ((i=0; i<n; i++)); do
        p=$((RANDOM % (126 - 33 + 1) + 33))
        input+=$(printf '%b' "\\$(printf '%03o' $p)")
    done
    
    output=$(echo "$input" | go run ../main.go)
    
    input="{\""$'\n'"$input"$'\n'"\"}"
    output="{\""$'\n'"$output"$'\n'"\"}"
    echo "$input" >> "$testSetFile"
    echo "$output" >> "$answerSetFile"
done

echo "Files written successfully"
