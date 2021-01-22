#!/bin/bash
echo "enter how many prints"
read argNum
declare -i integ=0
while [ $argNum -gt  $integ ]; do
	integ=$(($integ+1))
	printf "int: %i | string format %s\n" "$integ" "$integ"
done
