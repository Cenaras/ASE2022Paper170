#! /usr/bin/env bash

while read -r line
do
  for entry in $line
  do
    ./benchmarks/run.py --dir ${entry}_results --analysisStrategy ${entry}
  done
done < "./benchmarks/strategies.txt"
