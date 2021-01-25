package tempconv

//  Write a program that converts Centigrade to Fahrenheit
import "fmt"

func main() {
	//ftoC(5)
	fmt.Println("hej")
}

func FtoC(f int) int {
	var c int
	//fmt.Println("Enter C to convert it to fahrenheit.")
	// fmt.Scanln(&c)
	f = (c * 9 / 5) + 32
	fmt.Printf("%d°C  is in fahrenheit: %d°F \n", c, f)
	return f
}
