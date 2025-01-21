#!/bin/bash

testSetFile="test-set"
answerSetFile="answer-set"

touch "$testSetFile"  
touch "$answerSetFile"

for _ in {1..50}; do
    n=$((RANDOM * RANDOM % $((10**5)) + 1))
    input="$n"$'\n'
    for ((i = 0; i < n; i++)) do
        input+="$((RANDOM % 4 + 1)) "
    done
    input="${input%?}"
    
    output=$(echo "$input" | go run ../main.go)
    
    input="{\""$'\n'"$input"$'\n'"\"}"
    output="{\""$'\n'"$output"$'\n'"\"}"
    echo "$input" >> "$testSetFile"
    echo "$output" >> "$answerSetFile"
done

echo "Files written successfully"
