#!/bin/bash

testSetFile="test-set"
answerSetFile="answer-set"

touch "$testSetFile"  
touch "$answerSetFile"

for _ in {1..50}; do
    n=$((RANDOM % 100 + 1))
    input="$n"$'\n'
    shouldBeYes=$((RANDOM % 2))
    sum=(0 0 0)
    for ((i=0; i<n; i++)); do
        if ((i == n-1 && shouldBeYes == 1)); then
            input+="$((-sum[0])) $((-sum[1])) $((-sum[2]))"
            input+=$'\n'
            break
        fi
        for ((j=0; j<3; j++)); do
            num=$((RANDOM % $((200 + 1)) - 100))
            input+="$num "
            sum[j]=$((num + sum[j]))
        done
        input=${input::-1}
        input+=$'\n'
    done
    input=${input::-1}
    
    output=$(echo "$input" | go run ../main.go)
    
    input="{\""$'\n'"$input"$'\n'"\"}"
    output="{\""$'\n'"$output"$'\n'"\"}"
    echo "$input" >> "$testSetFile"
    echo "$output" >> "$answerSetFile"
done

echo "Files written successfully"
