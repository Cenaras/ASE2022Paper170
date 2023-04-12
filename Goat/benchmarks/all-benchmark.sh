#! /usr/bin/env bash

while read -r line
do
  for entry in $line
  do
    if [[ ${entry:0:1} != "#" ]]; then
      #echo $entry
      ./benchmarks/run.py --dir ${entry}_results --analysisStrategy ${entry}
    fi
  done
done < "./benchmarks/strategies.txt"
