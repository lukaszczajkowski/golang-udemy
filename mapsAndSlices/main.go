package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// don't do that usually
	//var myOtherMap map[string]string
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"

	myMap["dog"] = "fido"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	myNextMap := make(map[string]int)

	myNextMap["first"] = 1
	myNextMap["second"] = 2

	log.Println(myNextMap["first"], myNextMap["second"])

	userMap := make(map[string]User)

	me := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	userMap["me"] = me

	log.Println(userMap["me"].FirstName)

	// slices
	var mySlice []string
	mySlice = append(mySlice, "Anna")
	mySlice = append(mySlice, "Lukasz")
	mySlice = append(mySlice, "Yeti")

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)

	log.Println(numbers[0:2])
}
