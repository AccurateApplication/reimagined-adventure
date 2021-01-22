array=(one towoo thri)
arraystrings=("textfile" "testfile")
for i in $array; do
	echo $i
done
for i in $arraystrings; do
	echo $i
done
# not work


declare -a array2=("el1" "el2" "el3" "el4")
for i in "${array2[@]}";do
	echo " "
	echo $i
done

for item in hej tvo tre fyra; do
	printf "$item item\n"
done
