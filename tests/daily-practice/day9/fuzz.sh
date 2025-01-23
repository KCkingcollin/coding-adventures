#!/bin/bash

n=$((RANDOM % 100000 + 1))
input="$n"$'\n'
for ((i=0; i < n; i++)) do
    input+="$((RANDOM % 100000 + 1)) "
done
input=${input::-1}
q=$((RANDOM % 100000 + 1))
q=100000
input+=$'\n'"$q"$'\n'
for ((i=0; i < q; i++)) do
    input+="$((RANDOM % 110000 + 1))"$'\n'
done
input=${input::-1}

echo "$input"
