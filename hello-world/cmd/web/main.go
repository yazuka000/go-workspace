package main

import (
	"fmt"
	"myapp/pkg/handlers"
	"net/http"
)

var portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	// http.HandleFunc("/divide", Divide)

	fmt.Println((fmt.Sprintf("Starting application on port %s", portNumber)))
	_ = http.ListenAndServe(portNumber, nil)
}

// func addValues(x, y int) int {
// 	return x + y
// }

// func Divide(w http.ResponseWriter, r *http.Request) {
// 	f, err := divideValues(100.0, 0.0)
// 	if err != nil {
// 		fmt.Fprintf(w, "Cannot divide by 0")
// 		return
// 	}

// 	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
// }

// func divideValues(x, y float32) (float32, error) {
// 	if y <= 0 {
// 		err := errors.New("cannot divide by zero")
// 		return 0, err
// 	}
// 	result := x / y
// 	return result, nil
// }
