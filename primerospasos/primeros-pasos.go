package primerospasos

import "fmt"

func primerospasos() {
	//Practicas en goland
	fmt.Println("Pruebas de texto por pantalla")
	fmt.Println("Pruebas de texto por pantalla otra linea mas DOS")
	/*
		Este es un comentario de multiples linias
		Se suele usar para especificar como trabajan
		los distintos paquetes y sus funciones
	*/

	//variables
	var edad = 34
	var coche = "Mercedes"
	coche = "Toyota"

	// Con coma podemos concatenar valores
	fmt.Println("Mi edas es :", edad)

	// Print mas estricto
	fmt.Printf("%s %s", "Mi coche es un :", coche)
	fmt.Println("")
	// Entrada de datos
	var web string

	fmt.Print("¿Cual es tu sitio web?")
	fmt.Scan(&web) // Esta instruccion guarda en la variable web el valor que ingresa el usuario

	fmt.Println("Tu sitio web es:", web)

}
