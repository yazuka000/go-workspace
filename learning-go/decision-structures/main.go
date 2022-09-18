package main

import "log"

func main() {
	var isTrue1 bool
	isTrue1 = true

	if isTrue1 == true {
		log.Println("isTrue is", isTrue1)
	} else {
		log.Println("isTrue is", isTrue1)
	}

	cat := "cat"

	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	myNum := 100
	isTrue := false

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and isTrue is set to true")
	} else if myNum < 100 && isTrue {
		log.Println("1")
	}else if myNum == 101 || isTrue {
		log.Println("2")
	}else if myNum > 1000 && !isTrue {
		log.Println("3")
	}

	myVar := "horse"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")

	case "dog":
		log.Println("cat is set to dog")

	case "fish":
		log.Println("cat is set to fish")

	default:
		log.Println("cat is something else")
	}
}
