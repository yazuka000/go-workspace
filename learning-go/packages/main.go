package main

import (
	"github.com/yazuka000/myniceprogram/helpers"
	"log"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	log.Println(myVar.TypeName)
}
