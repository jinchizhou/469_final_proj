#!/bin/bash
# Basic until loop
counter=1
until [ $counter -gt 10 ]
do
random=$RANDOM
echo $random
time go run ring.go random
((counter++))
done
echo All done