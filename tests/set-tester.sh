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
    inputBlock="$testLine"
    expectedOutputBlock="$expectedLine"

    while true; do
        IFS= read -r testLine || break
        inputBlock+=$'\n'$testLine

        if [[ "$inputBlock" == *'{"'* && "$inputBlock" == *'"}'* ]]; then
            break
        fi
    done

    while true; do
        IFS= read -r expectedLine <&3 || break
        expectedOutputBlock+=$'\n'$expectedLine

        if [[ "$expectedOutputBlock" == *'{"'* && "$expectedOutputBlock" == *'"}'* ]]; then
            break
        fi
    done

    inputBlock="${inputBlock#*$'\n'}"
    inputBlock="${inputBlock%$'\n'*}"

    expectedOutputBlock="${expectedOutputBlock#*$'\n'}"
    expectedOutputBlock="${expectedOutputBlock%$'\n'*}"

    numLines=$(wc -l <<< "$inputBlock" | xargs)

    echo "Running test number $blockNum"
    actual_output=$($command <<< "$inputBlock")
    wait

    if [ "$actual_output" != "$expectedOutputBlock" ]; then
        echo "Output for block number '$blockNum' (starting at line number '$lineNum') does not match the expected output."
        echo "Expected: '$expectedOutputBlock'"
        echo "Got: '$actual_output'"
        echo "Input: '$inputBlock'"
        exit 1
    else
        echo "Passed"
    fi

    lineNum=$((lineNum + numLines))
    blockNum=$((blockNum + 1))
done < "$testSet" 3< "$answerSet"

echo "All outputs match"
exit 0
