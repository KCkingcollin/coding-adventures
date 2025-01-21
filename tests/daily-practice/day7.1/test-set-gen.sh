#!/bin/bash

testSetFile="test-set"
answerSetFile="answer-set"

touch "$testSetFile"  
touch "$answerSetFile"

function randomSting() {
    rnd="$(tr -dc '[:lower:]' < /dev/urandom | head -c "$1")"
    echo "$rnd"
}

for _ in {1..50}; do
    n=$((RANDOM % 100 + 1))
    input=""
    shouldBeYes=$((RANDOM % 2))
    str=(h e l l o)
    if ((shouldBeYes == 1)); then
        for char in "${str[@]}"; do
            input+="$(randomSting "$((n / 10))")$char$(randomSting "$((n / 10))")"
        done
    else
        input+="$(randomSting "$n")"
    fi
    
    output=$(echo "$input" | go run ../main.go)
    
    input="{\""$'\n'"$input"$'\n'"\"}"
    output="{\""$'\n'"$output"$'\n'"\"}"
    echo "$input" >> "$testSetFile"
    echo "$output" >> "$answerSetFile"
done

echo "Files written successfully"
