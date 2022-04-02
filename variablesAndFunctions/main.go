package main

import "fmt"

//package level variable
var myName string

func main() {
	fmt.Println("Hello, world.")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"

	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	watWasSaid, theOtherThingThatWasSaid := saySomething()

	fmt.Println("the function returned", watWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
