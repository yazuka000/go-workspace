package main

import (
	"log"
	"sort"
)

func main() {
	var stringSlice []string

	stringSlice = append(stringSlice, "Trevor")
	stringSlice = append(stringSlice, "John")
	stringSlice = append(stringSlice, "Mary")

	log.Println(stringSlice)

	var intSlice []int

	intSlice = append(intSlice, 2)
	intSlice = append(intSlice, 1)
	intSlice = append(intSlice, 3)

	sort.Ints(intSlice)

	log.Println(intSlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)

	log.Println(numbers[6:9])

	names := []string{"one", "seven", "fish", "cat"}

	log.Println(names)

}
