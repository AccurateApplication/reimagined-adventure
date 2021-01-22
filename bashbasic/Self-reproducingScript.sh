#!/bin/bash
# Write a script that backs itself up, that is, copies itself to a file named backup.sh.
# Hint: Use the cat command and the appropriate positional parameter.
command="cat ./textfile"
$command
FILE="textfile"
echo -e "\n"
command="cat $FILE"
$command "$0" > ScriptBackup
