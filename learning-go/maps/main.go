package main

import (
	"log"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	StrStrMap := make(map[string]string)

	StrStrMap["dog"] = "Samson"
	StrStrMap["other-dog"] = "Cassie"

	StrStrMap["dog"] = "fido"

	log.Println(StrStrMap["dog"])
	log.Println(StrStrMap["other-dog"])

	StrIntMap := make(map[string]int)

	StrIntMap["First"] = 1
	StrIntMap["Second"] = 2

	log.Println(StrIntMap["First"], StrIntMap["Second"])

	StrUserMap := make(map[string]User)

	me := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	StrUserMap["me"] = me

	log.Println(StrUserMap["me"].FirstName)

	var myNewVar float32

	myNewVar = 11.1

	log.Println(myNewVar)
}
