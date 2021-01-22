#!/bin/bash
echo -n "enter file: "
read FILE

if [[ -f $FILE ]] 
then
	echo "file exists"
else 
	echo "file does not exist"
fi


# something wrong below
day="date +%a"
$day
case $day in
	'Thu')
		echo "thursday"
		;;
	*) echo " case not work properly "
		;;
esac	
# 
