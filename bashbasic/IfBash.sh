#!/bin/bash
if test -e textfile; then echo "file exists";fi
if test -e textfil; then echo "file exists";else echo "file does not exist";fi
echo -e "\n"
file1=textfile
file2=textfil
if [ -e $file1 ]; then
	echo "[file1] exist"
else
	echo "[file1 does not exist]"
fi
if [ -e $file2 ]; then
	echo "[file2] exist"
else
	echo "[file2 does not exist]"
fi
exit 0
#man test
