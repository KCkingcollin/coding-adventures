#!/bin/bash

testSetFile="test-set"
answerSetFile="answer-set"

touch "$testSetFile"  
touch "$answerSetFile"

for _ in {1..50}; do
    n=$((RANDOM * RANDOM * RANDOM % $((10**9)) + 1))
    m=$((RANDOM * RANDOM * RANDOM % n + 1))
    a=$((RANDOM * RANDOM * RANDOM % n + 1))
    input="$n $m $a"
    
    output=$(echo "$input" | go run ../main.go)
    
    input="{\""$'\n'"$input"$'\n'"\"}"
    output="{\""$'\n'"$output"$'\n'"\"}"
    echo "$input" >> "$testSetFile"
    echo "$output" >> "$answerSetFile"
done

echo "Files written successfully"
