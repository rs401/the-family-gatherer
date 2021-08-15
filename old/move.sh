#!/usr/bin/bash

for file in *; do
    if [[ $file -eq "old" ]]; then
       continue
    fi
    mv $file old/$file
done
