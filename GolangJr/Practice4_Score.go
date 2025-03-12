package main

import "fmt"

func main() {
	var score float64

	fmt.Println("Enter your score: ")
	fmt.Scan(&score)

	if score >= 95.99 && score <= 100 {
		fmt.Println("Your value is A")
	} else if score >= 90.99 && score <= 95.99 {
		fmt.Println("Your value is B")
	} else if score >= 85.99 && score <= 90.99 {
		fmt.Println("Your value is C")
	} else if score >= 75.99 && score <= 85.99 {
		fmt.Println("Your value is D")
	} else if score >= 65.99 && score <= 75.99 {
		fmt.Println("Your value is E")
	} else {
		fmt.Println("Valor incorrecto")
	}

}
