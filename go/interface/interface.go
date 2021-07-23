package main

// https://research.swtch.com/interfaces
// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Define Animal as being any type that has method named Speak
// This takes no arguments and returns string
// Any type that defines this method "satifies" Animal interface.
type Animal interface {
	Speak() string
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

type Cat2 struct {
	name string
}

func (c Cat) Speak() string {
	return "Meow"
}

func (c *Cat2) Speak() string {
	return "2Meeeeow2"
}

func (d Dog) Speak() string {
	return "bark"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, new(Cat2), new(Dog)}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
	fmt.Println(Dog{}.Speak())
	moreInterface()

	names := []string{"Carl", "david", "Lemmi"}
	// Cannot use names (variable of type []string) as type []interface{} int arg to PrintAll
	// PrintAll(names)

	// Convert []string to []interface{}
	vals := make([]interface{}, len(names))
	for i, v := range names {
		vals[i] = v
	}
	PrintAll(vals)

	// our target will be of type map[string]interface{}, which is a pretty generic type
	// that will give us a hashtable whose keys are strings, and whose values are of
	// type interface{}
	var val map[string]interface{}

	if err := json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}

	fmt.Println(val)
	for k, v := range val {
		fmt.Println(k, reflect.TypeOf(v))
	}
}

var input = `
{
	"created_at": "Thu May 31 00:00:01 +0000 2012"
}
`

func moreInterface() {
	fmt.Println("tmp")

}

// Accepts any parameter
// v = type of interface{}. Not "any type"
func wowFunction(v interface{}) {
}
func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}

}

// An interface value is constructed of two words of data; one word is used to point to a method table for the value’s underlying type
// and the other word is used to point to the actual data being held by that value.
