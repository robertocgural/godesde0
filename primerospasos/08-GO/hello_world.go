package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hola, Go!")

	// Variables

	// variable string
	var myString string = "esto es una cadena de texto"
	fmt.Println(myString)

	myString = "Aqui cambio el valor de la cadena"
	fmt.Println(myString)

	// variable integer
	var myInt int = 7
	fmt.Println(myInt)

	myInt = myInt + 4
	fmt.Println(myInt)

	fmt.Println(myString, myInt)

	fmt.Println(reflect.TypeOf(myInt))

	// variable float
	var myFloat = 6.5
	fmt.Println(myFloat)
	fmt.Println(reflect.TypeOf(myFloat))

	fmt.Println(float64(myInt) + myFloat)

	// variable bool
	var myBool bool = false
	myBool = true
	fmt.Println(myBool)

	// Variable declarada e inicializada de manera abreviada
	myString2 := "Esto es una cadena de texto"
	fmt.Println(myString2)

	// Constantes
	const myConst = "Esto es una constante"
	fmt.Println(myConst)

	// Control de flujo (If - Else... etc)

}
