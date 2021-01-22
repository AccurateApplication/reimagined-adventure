#!/bin/bash
for i in "$@"; do
	echo "$i"
	echo "hej"
done
for i in "$@"; do
	printf "i: %s\n " "$i"
done
