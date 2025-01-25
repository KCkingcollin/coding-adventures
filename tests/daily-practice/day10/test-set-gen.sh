#!/bin/bash

testSetFile="test-set"
answerSetFile="answer-set"

touch "$testSetFile"  
touch "$answerSetFile"

for _ in {1..50}; do
    num=$((RANDOM % 100000 + 1))
    n=$((num / (RANDOM % 10 + 1)))
    q=$((num - n))
    input="$n"$'\n'
    for ((i=0; i < n; i++)) do
        input+="$((RANDOM % 100000 + 1)) "
    done
    input=${input::-1}
    input+=$'\n'"$q"$'\n'
    for ((i=0; i < q; i++)) do
        input+="$((RANDOM % 110000 + 1))"$'\n'
    done
    input=${input::-1}
    
    output=$(echo "$input" | go run ../main.go)
    
    input="{\""$'\n'"$input"$'\n'"\"}"
    output="{\""$'\n'"$output"$'\n'"\"}"
    echo "$input" >> "$testSetFile"
    echo "$output" >> "$answerSetFile"
done

echo "Files written successfully"
