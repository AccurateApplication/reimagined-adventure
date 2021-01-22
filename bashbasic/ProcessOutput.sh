#!/bin/bash
shelltest=$(ls | head -n 5)
echo $shelltest
echo -e "\n"
for i in $shelltest;do
	echo "$i for loop"
	#printf "percent T %T\n " "$i"
done
lslist=$(ls -l)
string=textfile
for i in $lslist ; do
	if [[ "$string" == "$i" ]] ;then
		echo "this file found: $i\t"
	fi
done
