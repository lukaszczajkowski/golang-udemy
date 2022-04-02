package main

import "log"

func main() {
	for i := 0; i <= 5; i++ {
		log.Println("i is set to", i)
	}

	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	for _, animal := range animals {
		log.Println(animal)
	}

	animalMap := make(map[string]string)

	animalMap["dog"] = "fido"
	animalMap["cat"] = "fluffy"

	for _, animal := range animalMap {
		log.Println("animal from map", animal)
	}

	for animalType, animal := range animalMap {
		log.Println(animalType, animal)
	}

	var firstLine = "Once upon a midnight dreary"

	for i, l := range firstLine {
		log.Println(i, ":", l)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User

	users = append(users, User{
		FirstName: "John",
		LastName:  "Smith",
		Email:     "john@smith.com",
		Age:       30,
	})
	users = append(users, User{
		FirstName: "Mary",
		LastName:  "Jones",
		Email:     "mary@jones.com",
		Age:       20,
	})
	users = append(users, User{
		FirstName: "Sally",
		LastName:  "Brown",
		Email:     "sally@smith.com",
		Age:       45,
	})
	users = append(users, User{
		FirstName: "Alex",
		LastName:  "Anderson",
		Email:     "alex@smith.com",
		Age:       17,
	})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
