package main

import (
	"github.com/lukaszczajkowski/myniceprogram/helpers"
	"log"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)

	intChan <- randomNumber
}

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	log.Println(myVar.TypeName)

	//channel
	intChan := make(chan int)
	// defer - whatever comes after defer, execute it as soon as the current function is done
	defer close(intChan)

	// routine - they run at the same time concurrently
	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
