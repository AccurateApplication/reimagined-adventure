#!/bin/bash
#for x in "$@" ; do
#	echo  $x
#
#done

#read hej
#NAME=$1
#if [ -z $NAME ]; then
#       echo enter name 
#	read NAME
#fi	
#echo $NAME

NAME=$1
NAME2=$2
if test -z $NAME ; then
       echo enter name 

	read NAME
fi	
echo $NAME2
echo $NAME
if [ -z $1 ]; then
       echo "enter name"       
       read NAME
else
	NAME=$1
fi
echo $NAME
exit 0
