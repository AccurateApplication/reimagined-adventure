package main

import "fmt"

func main() {
	//printcharactertype()
	//looponstring()
	//	countValley()
	//countingValleys(5, "UUUUUUU")
	countingValleys2(8, "UDDDUDUU")
}

func printcharactertype() {
	var stringz string = "hello mundo"

	fmt.Println("len function: ", len(stringz))
	for i := 0; i < len(stringz); i++ {
		fmt.Printf("%v \n", stringz[i])
		fmt.Println("stringz i", stringz[i])
		fmt.Printf("%T type .. .. . .", stringz[i])
	}
}

func looponstring() {
	var ord string = "UUDDDUUU"
	for index, char := range ord {
		fmt.Printf("character at index %d is %c\n", index, char)

	}
}

func countValley() {
	//	var n int32 = 5
	var s string = "UDDDUDUU"
	var lvl int = 1
	for n, s := range s {
		fmt.Printf("print n: %d and\t s: %c\n", n, s)
		fmt.Println("current lvl:\t", lvl)
		if s == 'U' {
			lvl++
		}
		if s == 'D' {
			lvl--
		}
	}
	fmt.Println("\n\n", lvl)
}

func countingValleys(n int, s string) int32 {
	var lvl int32 = 0
	fmt.Println("\ninnan loop")
	for i := 1; i <= n; i++ {
		fmt.Printf("LoopPrint i: %d \n", i)
		for x, z := range s {
			if z == 'U' {
				lvl++
			}
			if z == 'D' {
				lvl--
			}
			fmt.Println(x, lvl)
		}
	}
	fmt.Println("\n\tlvl: ", lvl, s)
	return lvl
}

func countingValleys1(n int, s string) int32 {
	var lvl int32 = 1
	//	var count int32 = 0
	for i := 0; i <= n; i++ {
		if s[i] == 'U' {
			lvl++
		}
		if s[i] == 'D' {
			lvl--
		}
	}
	//	(int i=0;i<n;i++){
	fmt.Println(lvl)
	return lvl
}

func countingValleys2(n int, s string) int32 {
	var lvl int32 = 0
	//	var count int32 = 0
	//	for i := 0; i <= n; i++ {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == 'U' {
			lvl++
		}
		if s[i] == 'D' {
			lvl--
		}
		if lvl == 0 && s[i] == 'D' {
			lvl++
		}
	}
	//	(int i=0;i<n;i++){
	fmt.Println(lvl)
	return lvl
}
