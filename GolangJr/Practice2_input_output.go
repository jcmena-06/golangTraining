package main

import (
	"fmt"
	"reflect"
)

func main() {

	var nombre string
	apellido := ""
	var edad int
	var estatura float64
	email := ""

	fmt.Println("Ingrese su nombre: ")
	fmt.Scanln(&nombre)
	fmt.Println("Ingrese su apellido: ")
	fmt.Scanln(&apellido)
	fmt.Println("Ingrese su edad: ")
	fmt.Scanln(&edad)
	fmt.Println("Ingrese su estatura en metros (Por ejemplo 1.75)")
	fmt.Scanln(&estatura)
	fmt.Println("Ingrese su email: ")
	fmt.Scanln(&email)

	//Imprimir los datos
	fmt.Println("\n📌 Datos ingresados:")
	fmt.Printf("Nombre: %s (Tipo: %s) \n", nombre, reflect.TypeOf(nombre))
	fmt.Printf("Apellido: %s (Tipo: %s) \n", apellido, reflect.TypeOf(apellido))
	fmt.Printf("Edad: %d (Tipo: %s)\n", edad, reflect.TypeOf(edad))
	fmt.Printf("Estatura: %.2f metros (Tipo: %s)\n", estatura, reflect.TypeOf(estatura))
	fmt.Printf("Su email es: %s (Tipo: %s) \n", email, reflect.TypeOf(email))
}
