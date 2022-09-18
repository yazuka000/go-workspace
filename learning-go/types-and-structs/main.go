package main

import (
	"log"
	"time"
)

// var firstName string
// var lastName string
// var phoneNumber string
// var age int
// var birthDate time.Time

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

// var Special string

func main() {
	user := User {
		FirstName: "Trevor",
		LastName: "Sawler",
		PhoneNumber: "1 555 555-1212",
	}

	log.Println(user.FirstName, user.LastName, "BirthDate:", user.BirthDate)
}

