#!/bin/bash

trap 'exit 1' ERR

if [ "$#" -ne 3 ]; then
    echo "Usage: $0 <testSet_location> <answerSet_location> <command>"
    exit 1
fi

testSet=$1
answerSet=$2
command=$3

if [ ! -f "$testSet" ]; then
    echo "Error: Test set file '$testSet' does not exist."
    exit 1
fi

if [ ! -f "$answerSet" ]; then
    echo "Error: Answer set file '$answerSet' does not exist."
    exit 1
fi

lineNum=1
blockNum=1
while read -r testLine && read -r expectedLine <&3; do
    inputBlock=""
    while true; do
        inputBlock+="$testLine"
        if echo "$inputBlock" | grep -q '{"' && echo "$inputBlock" | grep -q '"}'; then
            break
        fi
        inputBlock+=$'\n'
        IFS= read -r testLine || break
    done

    expectedOutputBlock=""
    while true; do
        expectedOutputBlock+="$expectedLine"
        if echo "$expectedOutputBlock" | grep -q '{"' && echo "$expectedOutputBlock" | grep -q '"}'; then
            break
        fi
        expectedOutputBlock+=$'\n'
        IFS= read -r expectedLine <&3 || break
    done

    inputBlock="${inputBlock:3:-3}"
    expectedOutputBlock="${expectedOutputBlock:3:-3}"

    numLines=$(echo -e "$inputBlock" | wc -l)

    echo "Running test number $blockNum"
    actual_output=$(echo -e "$inputBlock" | $command)
    wait

    if [ "$actual_output" != "$expectedOutputBlock" ]; then
        echo "Output for block number '$blockNum' (starting at line number '$lineNum') does not match the expected output."
        echo "Expected: '$expectedOutputBlock'"
        echo "Got: '$actual_output'"
        echo "Input: '$inputBlock'"
        exit 1
    fi

    lineNum=$((lineNum + numLines))
    blockNum=$((blockNum + 1))
done < "$testSet" 3< "$answerSet"

echo "All outputs match"
exit 0
