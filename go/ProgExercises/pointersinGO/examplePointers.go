package main

import "fmt"

func main() {
	x := 3459
	y := "hej fin string"
	fmt.Printf("Type of X: %T and Y: %T \n", x, y)
	fmt.Printf("Hex value of X: %X and Y: %X \n", x, y)
	fmt.Printf("Decimal value of X: %v and Y: %v \n", x, y)
	// address
	fmt.Println("Adress? of X: ", &x)
	z := 852858
	var p *int
	p = &x
	fmt.Println("this is printing \"p\" = Value stored in variable p = ", p)
	fmt.Println("USING &z returns ADDRESS OF Z: ", &z)
	fmt.Println("USING only Z returns :", z)

}
