#!/bin/bash

testSetFile="test-set"
answerSetFile="answer-set"

touch "$testSetFile"  
touch "$answerSetFile"

for _ in {1..50}; do
    n=$((RANDOM % 50 + 1))
    k=$((RANDOM % n + 1))
    input="$n $k"$'\n'
    for ((i=0; i<n; i++)); do
        a=$((RANDOM % 100))
        input+="$a "
    done
    input=${input::-1}
    
    output=$(echo "$input" | go run ../main.go)
    
    input="{\""$'\n'"$input"$'\n'"\"}"
    output="{\""$'\n'"$output"$'\n'"\"}"
    echo "$input" >> "$testSetFile"
    echo "$output" >> "$answerSetFile"
done

echo "Files written successfully"
