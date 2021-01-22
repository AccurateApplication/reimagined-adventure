#!/bin/bash
filename='textfile'
echo starting
while read i; do
	echo $i
	echo -e "\nnew rhhw"
done < $filename
echo -e "\nreverse:"
rev $filename
echo -e "\nupside down"
tac $filename
