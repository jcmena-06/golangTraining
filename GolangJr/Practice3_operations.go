package main

import "fmt"

func main() {
	var num1, num2 float64
	var operation string

	fmt.Print("Enter the first value: ")
	fmt.Scan(&num1)

	fmt.Println("Enter the second value: ")
	fmt.Scan(&num2)

	fmt.Println("Select one operation")
	fmt.Println("+. Suma")
	fmt.Println("-. Resta")
	fmt.Println("*. Multiplicacion")
	fmt.Println("/. División")
	fmt.Scan(&operation)

	var result float64

	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Error: no se puede dividir por cero.")
			return
		}
	default:
		fmt.Println("Operación no válida")
		return

	}
	fmt.Printf("El resultado de %.2f %s %.2f es: %.2f\n", num1, operation, num2, result)
}
