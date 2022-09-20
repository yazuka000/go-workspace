package main

import "log"

func main() {
	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for _, animal := range animals {
		log.Println(animal)
	}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	animals2 := make(map[string]string)
	animals2["dog"] = "Fido"
	animals2["cat"] = "Fluffy"

	for animalType, animal := range animals2 {
		log.Println(animalType, animal)
	}

	var firstLine = "Once upon a midnight dreary"
	firstLine = "x"

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
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@smith.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@smith.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@smith.com", 17})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
